package flagtypes

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/typ.v4"
)

// Output is a flag type for specifying the output format. E.g json vs yaml.
type Output string

// Possible output format values
const (
	OutputTree Output = "tree"
	OutputJSON Output = "json"
	OutputYAML Output = "yaml"
)

// ensures the type implements the interface
var _ pflag.Value = typ.Ref(Output(""))

// String returns the string representations for this flag.
func (o Output) String() string {
	return string(o)
}

// Set updates the existing flag's value, if parsing is successful.
func (o *Output) Set(new string) error {
	newOutput, err := ParseOutput(new)
	if err != nil {
		return err
	}
	*o = newOutput
	return nil
}

// ParseOutput will attempt to parse a given string as a valid output enum value.
func ParseOutput(s string) (Output, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case string(OutputTree):
		return OutputTree, nil
	case string(OutputJSON):
		return OutputJSON, nil
	case string(OutputYAML):
		return OutputYAML, nil
	default:
		return "", fmt.Errorf(`invalid output %q, must be one of "tree", "json", or "yaml"`, s)
	}
}

// Type returns this flag type's name.
func (Output) Type() string {
	return "output"
}
