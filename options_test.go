package guinea

import (
	"testing"
)

func TestOptionShortName(t *testing.T) {
	opt := Option{Name: "o"}
	if opt.String() != "-o" {
		t.Fatal("Invalid:", opt.String())
	}
}

func TestOptionLongName(t *testing.T) {
	opt := Option{Name: "option"}
	if opt.String() != "--option" {
		t.Fatal("Invalid:", opt.String())
	}
}

func TestArgumentSingularName(t *testing.T) {
	arg := Argument{Name: "argument", Multiple: false}
	if arg.String() != "<argument>" {
		t.Fatal("Invalid:", arg.String())
	}
}

func TestArgumentMultipleName(t *testing.T) {
	arg := Argument{Name: "argument", Multiple: true}
	if arg.String() != "<argument>..." {
		t.Fatal("Invalid:", arg.String())
	}
}

func TestArgumentOptionalName(t *testing.T) {
	arg := Argument{Name: "argument", Multiple: false, Optional: true}
	if arg.String() != "[<argument>]" {
		t.Fatal("Invalid:", arg.String())
	}
}

func TestArgumentMultipleOptionalName(t *testing.T) {
	arg := Argument{Name: "argument", Multiple: true, Optional: true}
	if arg.String() != "[<argument>...]" {
		t.Fatal("Invalid:", arg.String())
	}
}
