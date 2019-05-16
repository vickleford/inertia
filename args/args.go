package args

import (
	"errors"
	"flag"
	"fmt"
)

type Config struct {
	Vars     Inputs
	Template string
}

func Parse(cmdline []string) (*Config, error) {
	c := new(Config)
	sliceOfVars := make(Inputs)
	sliceOfB64Vars := make(Base64Inputs)
	fs := flag.NewFlagSet("inertia", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(),
			"Usage: inertia [ -set templatekey=value ... ] [ -b64set templatekey=value ... ] /path/to/template\n")
		fs.PrintDefaults()
	}

	fs.Var(&sliceOfVars, "set", "Set templatekey with value")
	fs.Var(&sliceOfB64Vars, "b64set", "Set templatekey with the base64 encoding of value")

	err := fs.Parse(cmdline)
	if err != nil {
		return nil, err
	}

	c.Vars = sliceOfVars

	for k, v := range sliceOfB64Vars {
		c.Vars[k] = v
	}

	if leftovers := fs.Args(); len(leftovers) > 0 {
		c.Template = fs.Args()[0]
	} else {
		fs.Usage()
		return nil, errors.New("Incorrect usage")
	}

	return c, nil
}
