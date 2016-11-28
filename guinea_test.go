package guinea

import (
	"testing"
)

func TestFindCommand(t *testing.T) {
	var subSubCmdA = Command{}

	var subSubCmdB = Command{}

	var subCmdA = Command{
		Subcommands: map[string]*Command{
			"subA": &subSubCmdA,
			"subB": &subSubCmdB,
		},
	}

	var subCmdB = Command{}

	var mainCmd = Command{
		Subcommands: map[string]*Command{
			"subA": &subCmdA,
			"subB": &subCmdB,
		},
	}

	// mainCmd --- subA --- subA
	//          |_ subB  |_ subB

	cmd, cmdName, cmdArgs := FindCommand(&mainCmd, []string{"program"})
	if cmd != &mainCmd {
		t.Fatal("Invalid cmd")
	}
	if cmdName != "program" {
		t.Fatal("Invalid cmdName")
	}
	if len(cmdArgs) != 0 {
		t.Fatal("Invalid cmdArgs")
	}

	cmd, cmdName, cmdArgs = FindCommand(&mainCmd, []string{"program", "subA", "subB", "param1", "param2"})
	if cmd != &subSubCmdB {
		t.Fatal("Invalid cmd")
	}
	if cmdName != "program subA subB" {
		t.Fatal("Invalid cmdName")
	}
	if len(cmdArgs) != 2 {
		t.Fatal("Invalid cmdArgs")
	}
}
