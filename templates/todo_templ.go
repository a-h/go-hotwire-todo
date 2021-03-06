// Code generated by templ DO NOT EDIT.

package templates

import "github.com/a-h/templ"
import "context"
import "io"
import todo "github.com/a-h/go-hotwire-todo"

func Page(title string, content templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<html>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(title))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<script")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " src=\"https://unpkg.com/@hotwired/turbo@7.0.0-beta.5/dist/turbo.es5-umd.js\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</script>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<body>")
		if err != nil {
			return err
		}
		err = content.Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</html>")
		if err != nil {
			return err
		}
		return err
	})
}

func Index(d IndexViewData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Todos"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		err = Todos(d.Todos).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Create"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		err = NewTodo(d.NewTodo).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<turbo-frame")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=\"remote-frame\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " src=\"http://localhost:8001/remote-frame\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("Loading from remote..."))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</turbo-frame>")
		if err != nil {
			return err
		}
		return err
	})
}

func Todos(todos []*todo.Todo) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<turbo-frame")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=\"todos\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		for _, t := range todos {
			err = Todo(t).Render(ctx, w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "</turbo-frame>")
		if err != nil {
			return err
		}
		return err
	})
}

func Todo(t *todo.Todo) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(t.Item))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func NewTodo(d NewTodoViewData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<turbo-frame")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=\"new_todo\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<form")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=\"/\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " method=\"post\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"text\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"text\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(d.Text))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</input>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		if d.TextValidation != "" {
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color: red\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, templ.EscapeString(d.TextValidation))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=\"New\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</input>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</turbo-frame>")
		if err != nil {
			return err
		}
		return err
	})
}

