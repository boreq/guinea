package guinea

import (
	"testing"
)

func TestCommandWrongArguments(t *testing.T) {
	var mainCmd = Command{
		Arguments: []Argument{
			{Name: "arg1"},
			{Name: "arg2"},
		},
	}

	if mainCmd.Execute("program", []string{"a"}) != ErrInvalidParms {
		t.Fatal("Execute did not return ErrInvalidParams")
	}
}
