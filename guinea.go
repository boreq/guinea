package guinea

import (
	"os"
	"strings"
)

var globalOpt = []Option{
	Option{
		Name:        "help",
		Type:        Bool,
		Default:     false,
		Description: "Display help",
	},
}

// Run is a high level function which adds special behaviour to the commands,
// namely displaying help to the user. If you wish to use the library without
// that feature use the FindCommand function directly.
func Run(rootCommand *Command) error {
	cmd, cmdName, cmdArgs := FindCommand(rootCommand, os.Args)
	cmd.Options = append(cmd.Options, globalOpt...)
	return cmd.Execute(cmdName, cmdArgs)
}

// FindCommand attempts to recursively locate the command which should be
// executed. The provided command should be the root command of the program
// containing all other subcommands. The array containing the provided
// arguments will most likely be the os.Args array. The function returns the
// located subcommand, its name and the remaining unused arguments. Those
// values should be passed to the Command.Execute method.
func FindCommand(cmd *Command, args []string) (*Command, string, []string) {
	foundCmd, foundArgs := findCommand(cmd, args[1:])
	foundName := subcommandName(args, foundArgs)
	return foundCmd, foundName, foundArgs
}

func findCommand(cmd *Command, args []string) (*Command, []string) {
	for subCmdName, subCmd := range cmd.Subcommands {
		if len(args) > 0 && args[0] == subCmdName {
			return findCommand(subCmd, args[1:])
		}
	}
	return cmd, args
}

func subcommandName(originalArgs []string, remainingArgs []string) string {
	argOffset := len(originalArgs) - len(remainingArgs)
	return strings.Join(originalArgs[:argOffset], " ")
}
