package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	todo "github.com/a-h/go-hotwire-todo"
	"github.com/a-h/go-hotwire-todo/templates"
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

func main() {
	// Wire up the routes.
	http.Handle("/", IndexHandler{})
	// Use localhost:8000 so MacOS doesn't bug you to ask if you want to accept incoming connections.
	fmt.Println("Listening on http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// TODO: Add websockets?
	// https://discuss.hotwire.dev/t/how-to-connect-turbo-streams-to-spring-boot-websocket/1689/3
}

// Turbo stuff. This would be moved to a library.
type ActionType string

const (
	ActionAppend  ActionType = "append"
	ActionPrepend            = "prepend"
	ActionReplace            = "replace"
	ActionUpdate             = "update"
	ActionRemove             = "remove"
)

type Action struct {
	Type     ActionType
	Target   string
	Template templ.Component
}

func StreamAction(at ActionType, target string, template templ.Component) Action {
	return Action{Type: at, Target: target, Template: template}
}

func TurboStream(actions ...Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/vnd.turbo-stream.html")
		for _, action := range actions {
			action := action
			templates.Action(string(action.Type), action.Target, action.Template).Render(r.Context(), w)
		}
	})
}

func IsTurboRequest(r *http.Request) bool {
	return strings.Contains(r.Header.Get("accept"), "text/vnd.turbo-stream.html")
}

// End of turbo stuff.

// Create a decoder to handle getting values from form posts and populating structs.
// https://github.com/gorilla/schema
var decoder = schema.NewDecoder()

type IndexHandler struct{}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
		return
	case http.MethodPost:
		h.Post(w, r)
		return
	}
	http.Error(w, "unhandled verb", http.StatusBadRequest)
}

func (h IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	todos, err := todo.DB{}.List()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	d := templates.IndexViewData{
		Todos:   todos,
		NewTodo: templates.NewTodoViewData{},
	}
	h.Render(d, w, r)
}

func (h IndexHandler) Render(d templates.IndexViewData, w http.ResponseWriter, r *http.Request) {
	// Create the templates.
	body := templates.Index(d)
	page := templates.Page("Todos", body)

	// Render.
	err := page.Render(r.Context(), w)
	if err != nil {
		log.Println("error", err)
	}
}

func (h IndexHandler) Post(w http.ResponseWriter, r *http.Request) {
	// Parse the form.
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form post", http.StatusBadRequest)
		return
	}

	// Populate the structs.
	ntvd := new(templates.NewTodoViewData)
	err = decoder.Decode(ntvd, r.PostForm)
	if err != nil {
		http.Error(w, "failed to decode form post", http.StatusBadRequest)
		return
	}

	// Validate and carry out actions.
	isValid := ntvd.Validate()
	var newTodo *todo.Todo
	if isValid {
		// Update the data.
		newTodo = &todo.Todo{
			ID:   uuid.New().String(),
			Item: ntvd.Text,
		}
		todo.DB{}.Upsert(newTodo.ID, newTodo.Item, newTodo.Complete)
		// Clear the form.
		ntvd = new(templates.NewTodoViewData)
	}

	// Get the view ready.
	// If it's not a Turbo request, we have to render the whole screen. However, this means
	// the app works without JavaScript.
	if !IsTurboRequest(r) {
		todos, err := todo.DB{}.List()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		d := templates.IndexViewData{
			Todos:   todos,
			NewTodo: *ntvd,
		}
		h.Render(d, w, r)
		return
	}

	// If it's a Turbo request, we can just update the bits of the screen we need to.
	var actions []Action

	// Update the todo list.
	if newTodo != nil {
		actions = append(actions, StreamAction(ActionAppend, "todos", templates.Todo(newTodo)))
	}

	// Update the form.
	actions = append(actions, StreamAction(ActionUpdate, "new_todo", templates.NewTodo(*ntvd)))

	// Return the stream of updates.
	TurboStream(actions...).ServeHTTP(w, r)
}
