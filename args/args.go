package args

import (
	"flag"
	"fmt"
	"strings"
)

type vars map[string]string

func (v *vars) String() string {
	var b strings.Builder
	for k, val := range *v {
		s := fmt.Sprintf("-set %s=%s ", k, val)
		b.WriteString(s)
	}
	return b.String()
}

func (v *vars) Set(value string) error {
	parts := strings.Split(value, "=")
	(*v)[parts[0]] = strings.Join(parts[1:], "")
	return nil
}

type Config struct {
	Vars     *vars
	Image    string
	Template string
}

func Parse(cmdline []string) (*Config, error) {
	c := new(Config)
	sliceOfVars := make(vars)
	fs := flag.NewFlagSet("Whatever", flag.ExitOnError)

	fs.Var(&sliceOfVars, "set", "Do this usage with this")

	err := fs.Parse(cmdline)
	if err != nil {
		fmt.Printf("omg got an error: %s", err)
		return nil, err
	}
	c.Vars = &sliceOfVars

	c.Template = fs.Args()[0]
	c.Image = fs.Args()[1]

	return c, nil
}
