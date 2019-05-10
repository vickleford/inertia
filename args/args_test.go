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

func TestSetWithTwoEqualSigns(t *testing.T) {
	mock := []string{"-set", "bizz=buz==", "template"}
	config, err := args.Parse(mock)
	if err != nil {
		t.Errorf("burn, got error: %s", err)
	}

	if actual := config.Vars["bizz"]; actual != "buz==" {
		t.Errorf("Wanted 'buz==' but got '%s'", actual)
	}
}

func TestSetRepeatingB64Args(t *testing.T) {
	const b64foobar = "Zm9vYmFy"
	const b64bazonii = "YmF6b25paQ=="
	const b64dennisbrown = "ZGVubmlzYnJvd24="

	mock := []string{
		"-b64set", "one=foobar",
		"-b64set", "two=bazonii", 
		"-b64set", "bestartist=dennisbrown", 
		"template",
	}
	config, err := args.Parse(mock)
	if err != nil {
		t.Errorf("burn, got error: %s", err)
	}

	if actual := config.Vars["one"]; actual != b64foobar {
		t.Errorf("Wanted '%s' but got '%s'", b64foobar, actual)
	}

	if actual := config.Vars["two"]; actual != b64bazonii {
		t.Errorf("Wanted '%s' but got '%s'", b64bazonii, actual)
	}

	if actual := config.Vars["bestartist"]; actual != b64dennisbrown {
		t.Errorf("Wanted '%s' but got '%s'", b64dennisbrown, actual)
	}
}

func TestTemplateGetsSet(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template"}
	config, _ := args.Parse(mock)

	if actual := config.Template; actual != "template" {
		t.Errorf("Wanted 'template' but got '%s'", actual)
	}
}

func TestInputsToString(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template"}
	expected := "-set foo=bar -set baz=buz "
	alternative := "-set baz=buz -set foo=bar "
	config, _ := args.Parse(mock)
	if actual := config.Vars.String(); !(actual == expected || actual == alternative) {
		t.Errorf("Wanted '%s' or '%s' but got '%s'", expected, alternative, actual)
	}
}

func TestBase64InputsToString(t *testing.T) {
	expected := "-b64set foo=YmFy "
	b64inputs := make(args.Base64Inputs)
	b64inputs.Set("foo=bar")
	if actual := b64inputs.String(); actual != expected {
		t.Errorf("Wanted '%s' but got '%s'", expected, actual)	
	}
}

// think about: cat template | inertia -set foo=bar | kubectl apply -f -

// need an error when template is presented before the arguments.
// aluminum13:inertia vwatkins$ go run main.go rendering/testdata/simple.tpl -set image=foo -set adjective=hairy -set noun=bumpkin
// <no value> has a big ole <no value> <no value> insidealuminum13:inertia vwatkins$
// aluminum13:inertia vwatkins$ go run main.go -set image=foo -set adjective=hairy -set noun=bumpkin rendering/testdata/simple.tpl
// foo has a big ole hairy bumpkin insidealuminum13:inertia vwatkins$
