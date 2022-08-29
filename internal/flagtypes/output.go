package flagtypes

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/typ.v4"
)

type Output string

const (
	OutputTree Output = "tree"
	OutputJSON Output = "json"
	OutputYAML Output = "yaml"
)

var _ pflag.Value = typ.Ref(Output(""))

func (o Output) String() string {
	return string(o)
}

func (o *Output) Set(new string) error {
	newOutput, err := ParseOutput(new)
	if err != nil {
		return err
	}
	*o = newOutput
	return nil
}

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

func (o Output) Type() string {
	return "output"
}
