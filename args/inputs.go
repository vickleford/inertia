package args

import (
	"fmt"
	"strings"
	"encoding/base64"
)

type Inputs map[string]string

func (v *Inputs) String() string {
	return toString(*v, "-set")}

func (v *Inputs) Set(value string) error {
	key, val := split(value)
	(*v)[key] = val
	return nil
}

type Base64Inputs map[string]string

func (v *Base64Inputs) String() string {
	return toString(*v, "-b64set")
}

func (v *Base64Inputs) Set(value string) error {
	key, val := split(value)
	b64val := base64.StdEncoding.EncodeToString([]byte(val))
	(*v)[key] = b64val
	return nil
}

func split(value string) (string, string) {
	const divider = "="
	dividerIndex := strings.Index(value, divider)
	key := value[0:dividerIndex]
	val := value[dividerIndex+1:]
	return key, val
}

func toString(i map[string]string, flag string) string {
	var b strings.Builder
	for k, v := range i {
		s := fmt.Sprintf("%s %s=%s ", flag, k, v)
		b.WriteString(s)
	}
	return b.String()
}
