package args

import (
	"flag"
	"fmt"
	"strings"
	"encoding/base64"
)

const divider = "="

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
	dividerIndex := strings.Index(value, divider)
	key := value[0:dividerIndex]
	val := value[dividerIndex+1:]
	(*v)[key] = val
	return nil
}

type Base64Inputs map[string]string

func (v *Base64Inputs) String() string {
	var b strings.Builder
	for k, val := range *v {
		s := fmt.Sprintf("-b64set %s=%s ", k, val)
		b.WriteString(s)
	}
	return b.String()
}

func (v *Base64Inputs) Set(value string) error {
	dividerIndex := strings.Index(value, divider)
	key := value[0:dividerIndex]
	val := value[dividerIndex+1:]
	b64val := base64.StdEncoding.EncodeToString([]byte(val))
	(*v)[key] = b64val
	return nil
}

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
		fmt.Printf("omg got an error: %s", err)
		return nil, err
	}
	c.Vars = sliceOfVars

	for k, v := range sliceOfB64Vars {
		c.Vars[k] = v
	}

	c.Template = fs.Args()[0]

	return c, nil
}
