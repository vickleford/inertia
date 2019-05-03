package rendering_test

import (
	"strings"
	"testing"

	"github.com/vickleford/inertia/args"
	"github.com/vickleford/inertia/rendering"
)

func TestRender(t *testing.T) {
	expected := "foo/bar:tag has a big ole hairy wumplewomper inside"
	output := new(strings.Builder)

	inputs := make(args.Inputs)
	inputs["image"] = "foo/bar:tag"
	inputs["adjective"] = "hairy"
	inputs["noun"] = "wumplewomper"

	// templateBytes, err := ioutil.ReadFile("testdata/simple.tpl")
	// if err != nil {
	// 	t.Errorf("Error reading test data: %s", err)
	// }

	tpl := "{{ .image }} has a big ole {{ .adjective}} {{ .noun }} inside"

	err := rendering.Render(tpl, inputs, output)
	if err != nil {
		t.Errorf("Rendering blew up: %s", err)
	}
	if output.String() != expected {
		t.Errorf("Wanted %s but got %s", expected, output.String())
	}
}
