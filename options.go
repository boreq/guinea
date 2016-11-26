package cli

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

func (arg Argument) String() string {
	format := "<%s>"
	if arg.Multiple {
		format += "..."
	}
	return fmt.Sprintf(format, arg.Name)
}
