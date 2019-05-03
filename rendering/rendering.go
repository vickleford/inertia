package rendering

import (
	"github.com/vickleford/inertia/args"
	"io"
	"text/template"
)

func Render(source string, in args.Inputs, output io.Writer) error {
	tmpl, err := template.New("render").Parse(source)
	if err != nil {
		return err
	}
	err = tmpl.Execute(output, in)
	if err != nil {
		return err
	}

	return nil
}
