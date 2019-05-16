package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vickleford/inertia/args"
	"github.com/vickleford/inertia/rendering"
)

func main() {
	cfg, err := args.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	tf, err := openSesame(cfg.Template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %s", cfg.Template, err)
		os.Exit(1)
	}

	tplbytes, err := ioutil.ReadAll(tf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading template: %s", err)
		os.Exit(1)
	}

	err = rendering.Render(string(tplbytes), cfg.Vars, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering: %s", err)
		os.Exit(1)
	}
}

func openSesame(filename string) (*os.File, error) {
	var f *os.File
	if filename == "-" {
		return os.Stdin, nil
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}
