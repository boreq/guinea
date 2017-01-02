package guinea

import (
	"fmt"
)

type ValType int

const (
	String ValType = iota
	Bool
	Int
)

// Option represents an optional flag.
type Option struct {
	Name        string
	Type        ValType
	Default     interface{}
	Description string
}

// String prepends the option name with one or two leading dashes and returns
// it. It is used to generate help texts.
func (opt Option) String() string {
	prefix := "-"
	if len(opt.Name) > 1 {
		prefix = "--"
	}
	return prefix + opt.Name
}

// Argument represents a required argument.
type Argument struct {
	Name        string
	Multiple    bool
	Description string
}

// String places the argument name in angle brackets and appends three dots to
// it in order to indicate multiple arguments. It is used to generate help
// texts.
func (arg Argument) String() string {
	format := "<%s>"
	if arg.Multiple {
		format += "..."
	}
	return fmt.Sprintf(format, arg.Name)
}
