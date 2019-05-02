package rendering_test

import (
	"os"
	"testing"
	"text/template"
)

func TestIfCanUseMap(t *testing.T) {
	vars := make(map[string]string)
	vars["image"] = "foo/bar:tag"
	tmpl, err := template.New("test").Parse("thing has {{ .image }} inside")
	if err != nil {
		t.Errorf("bomb: %s", err)
	}
	err = tmpl.Execute(os.Stdout, vars)
	if err != nil {
		t.Errorf("bomb: %s", err)
	}
	t.Errorf("let me see stdout now")
}
