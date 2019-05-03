package args

import (
	"flag"
	"fmt"
	"strings"
)

type Inputs map[string]string

func (v *Inputs) String() string {
	var b strings.Builder
	for k, val := range *v {
		s := fmt.Sprintf("-set %s=%s ", k, val)
		b.WriteString(s)
	}
	return b.String()
}

func (v *Inputs) Set(value string) error {
	parts := strings.Split(value, "=")
	(*v)[parts[0]] = strings.Join(parts[1:], "")
	return nil
}

type Config struct {
	Vars     Inputs
	Template string
}

func Parse(cmdline []string) (*Config, error) {
	c := new(Config)
	sliceOfVars := make(Inputs)
	fs := flag.NewFlagSet("Whatever", flag.ExitOnError)

	fs.Var(&sliceOfVars, "set", "Do this usage with this")

	err := fs.Parse(cmdline)
	if err != nil {
		fmt.Printf("omg got an error: %s", err)
		return nil, err
	}
	c.Vars = sliceOfVars

	c.Template = fs.Args()[0]

	return c, nil
}
