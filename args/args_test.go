package args_test

import (
	"github.com/vickleford/inertia/args"
	"testing"
)

func TestSetRepeatingArgs(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template"}

	config, err := args.Parse(mock)
	if err != nil {
		t.Errorf("burn, got error: %s", err)
	}

	if actual := config.Vars["foo"]; actual != "bar" {
		t.Errorf("Wanted 'bar' but got '%s'", actual)
	}

	if actual := config.Vars["baz"]; actual != "buz" {
		t.Errorf("Wanted 'buz' but got '%s'", actual)
	}
}

func TestSetWithTwoEqualSisng(t *testing.T) {
	t.Skip("Come back to this later")
	mock := []string{"-set", "bizz=buz=="}
	config, err := args.Parse(mock)
	if err != nil {
		t.Errorf("burn, got error: %s", err)
	}

	if actual := config.Vars["bizz"]; actual != "buz==" {
		t.Errorf("Wanted 'buz==' but got '%s'", actual)
	}
}

func TestTemplateGetsSet(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template"}
	config, _ := args.Parse(mock)

	if actual := config.Template; actual != "template" {
		t.Errorf("Wanted 'template' but got '%s'", actual)
	}
}

func TestVarsToString(t *testing.T) {
	expected := "-set foo=bar -set baz=buz "
	alternative := "-set baz=buz -set foo=bar "
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template"}
	config, _ := args.Parse(mock)
	if actual := config.Vars.String(); !(actual == expected || actual == alternative) {
		t.Errorf("Wanted '%s' or '%s' but got '%s'", expected, alternative, actual)
	}
}

// think about: cat template | inertia -set foo=bar | kubectl apply -f -

// need an error when template is presented before the arguments.
// aluminum13:inertia vwatkins$ go run main.go rendering/testdata/simple.tpl -set image=foo -set adjective=hairy -set noun=bumpkin
// <no value> has a big ole <no value> <no value> insidealuminum13:inertia vwatkins$
// aluminum13:inertia vwatkins$ go run main.go -set image=foo -set adjective=hairy -set noun=bumpkin rendering/testdata/simple.tpl
// foo has a big ole hairy bumpkin insidealuminum13:inertia vwatkins$
