package guinea

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidParms can be returned by a CommandFunction to automatically
// display help text.
var ErrInvalidParms = errors.New("invalid parameters")

type CommandFunction func(Context) error

// Command represents a single command which can be executed by the program.
type Command struct {
	Run              CommandFunction
	Subcommands      map[string]*Command
	Options          []Option
	Arguments        []Argument
	ShortDescription string
	Description      string
}

// UsageString returns a short string containing the syntax of this command.
func (c Command) UsageString(cmdName string) string {
	rw := cmdName
	if len(c.Subcommands) > 0 {
		rw += " <subcommand>"
	}
	rw += " [<options>]"
	for _, arg := range c.Arguments {
		rw += fmt.Sprintf(" %s", arg)
	}
	return rw
}

// PrintHelp prints the full help text to the standard output. The text
// contains the syntax of the command, a description, a list of accepted
// options and arguments and available subcommands.
func (c Command) PrintHelp(cmdName string) {
	usage := c.UsageString(cmdName)
	fmt.Printf("\n    %s - %s\n", usage, c.ShortDescription)

	if len(c.Options) > 0 {
		fmt.Println("\nOPTIONS:")
		for _, opt := range c.Options {
			fmt.Printf("    %-20s %s\n", opt, opt.Description)
		}
	}

	if len(c.Arguments) > 0 {
		fmt.Println("\nARGUMENTS:")
		for _, arg := range c.Arguments {
			fmt.Printf("    %-20s %s\n", arg, arg.Description)
		}
	}

	if len(c.Subcommands) > 0 {
		fmt.Println("\nSUBCOMMANDS:")
		for name, subCmd := range c.Subcommands {
			fmt.Printf("    %-20s %s\n", name, subCmd.ShortDescription)
		}
		fmt.Printf("\n    Try '%s <subcommand> --help'\n", cmdName)
	}

	if len(c.Description) > 0 {
		fmt.Println("\nDESCRIPTION:")
		desc := strings.Trim(c.Description, "\n")
		for _, line := range strings.Split(desc, "\n") {
			fmt.Printf("    %s\n", line)
		}
	}
}

// Execute runs the command. Command name is used to generate the help texts
// and should usually be set to one of the return values of FindCommand. The
// array of the arguments provided for this subcommand is used to generate the
// context and should be set to one of the return values of FindCommand as
// well. The command will not be executed with an insufficient number of
// arguments so there is no need to check that in the run function.
func (c Command) Execute(cmdName string, cmdArgs []string) error {
	context, err := makeContext(c, cmdArgs)
	if err != nil {
		c.PrintHelp(cmdName)
		return err
	}

	// Is the number of arguments sufficient?
	if err := c.validateArgs(context.Arguments); err != nil {
		c.PrintHelp(cmdName)
		return err
	}

	// Is there a help flag and is it set?
	if help, ok := context.Options["help"]; ok && help.Bool() {
		c.PrintHelp(cmdName)
		return nil
	}

	// Is this command only used to hold subcommands?
	if c.Run == nil {
		c.PrintHelp(cmdName)
		return nil
	}

	e := c.Run(*context)
	if e == ErrInvalidParms {
		c.PrintHelp(cmdName)
	}
	return e
}

func (c Command) validateArgs(cmdArgs []string) error {
	if len(cmdArgs) < len(c.Arguments) {
		return ErrInvalidParms
	}
	return nil
}
