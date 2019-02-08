package guinea

import (
	"os"
	"testing"
)

// Assigns /dev/null to stdout and returns a function for restoring it to the original one.
func supressStdout() func() {
	stdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	return func() {
		os.Stdout = stdout
	}
}

func TestCommandWrongArguments(t *testing.T) {
	var mainCmd = Command{
		Arguments: []Argument{
			{Name: "arg1"},
			{Name: "arg2"},
		},
	}

	restoreStdout := supressStdout()
	defer restoreStdout()

	if mainCmd.Execute("program", []string{"a"}) != ErrInvalidParms {
		t.Fatal("Execute did not return ErrInvalidParams")
	}
}

func TestExecuteCommandWithArgsPassingHelpShouldntReturnError(t *testing.T) {
	var cmd = Command{
		Options: []Option{
			{
				Name:        "help",
				Type:        Bool,
				Default:     false,
				Description: "Display help",
			},
		},
		Arguments: []Argument{
			{Name: "arg"},
		},
	}

	restoreStdout := supressStdout()
	defer restoreStdout()

	if err := cmd.Execute("prog sub", []string{"--help"}); err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}
