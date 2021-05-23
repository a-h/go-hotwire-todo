package templates

import todo "github.com/a-h/go-hotwire-todo"

// The index page.
type IndexViewData struct {
	Todos   []*todo.Todo
	NewTodo NewTodoViewData
}
type NewTodoViewData struct {
	Text string `schema:"text"`
	// This can be done with a required attribute in HTML, there's no need for this, it's just an example.
	// https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation
	TextValidation string
}

func (d *NewTodoViewData) Validate() (isValid bool) {
	isValid = true
	if d.Text == "" {
		d.TextValidation = "Text cannot be empty."
		isValid = false
	}
	return
}
