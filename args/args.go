package args

import (
	"flag"
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

	fs.Var(&sliceOfVars, "set", "Do this usage with this")
	fs.Var(&sliceOfB64Vars, "b64set", "Do this usage with this")

	err := fs.Parse(cmdline)
	if err != nil {
		return nil, err
	}

	c.Vars = sliceOfVars

	for k, v := range sliceOfB64Vars {
		c.Vars[k] = v
	}

	c.Template = fs.Args()[0]

	return c, nil
}
