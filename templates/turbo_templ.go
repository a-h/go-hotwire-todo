// Code generated by templ DO NOT EDIT.

package templates

import "github.com/a-h/templ"
import "context"
import "io"

func Action(action string, target string, template templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<turbo-stream")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(action))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " target=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(target))
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
		_, err = io.WriteString(w, "<template>")
		if err != nil {
			return err
		}
		err = template.Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</template>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</turbo-stream>")
		if err != nil {
			return err
		}
		return err
	})
}
