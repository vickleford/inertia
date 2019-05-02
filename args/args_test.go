package args_test

import (
	"github.com/vickleford/inertia/args"
	"testing"
)

func TestSetRepeatingArgs(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template", "image"}

	config, err := args.Parse(mock)
	if err != nil {
		t.Errorf("burn, got error: %s", err)
	}

	if actual := (*config.Vars)["foo"]; actual != "bar" {
		t.Errorf("Wanted 'bar' but got '%s'", actual)
	}

	if actual := (*config.Vars)["baz"]; actual != "buz" {
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

	if actual := (*config.Vars)["bizz"]; actual != "buz==" {
		t.Errorf("Wanted 'buz==' but got '%s'", actual)
	}
}

func TestImageGetsSet(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template", "ecrurl/image:tag"}
	config, _ := args.Parse(mock)

	if actual := config.Image; actual != "ecrurl/image:tag" {
		t.Errorf("Wanted 'ecrurl/iagme:tag' but got '%s'", actual)
	}
}

func TestTemplateGetsSet(t *testing.T) {
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template", "image"}
	config, _ := args.Parse(mock)

	if actual := config.Template; actual != "template" {
		t.Errorf("Wanted 'template' but got '%s'", actual)
	}
}

func TestVarsToString(t *testing.T) {
	expected := "-set foo=bar -set baz=buz "
	alternative := "-set baz=buz -set foo=bar "
	mock := []string{"-set", "foo=bar", "-set", "baz=buz", "template", "image"}
	config, _ := args.Parse(mock)
	if actual := config.Vars.String(); !(actual == expected || actual == alternative) {
		t.Errorf("Wanted '%s' or '%s' but got '%s'", expected, alternative, actual)
	}
}
