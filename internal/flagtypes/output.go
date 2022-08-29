// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

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
