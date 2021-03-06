package activityhelpers

import (
	"testing"
)

func TestIsOnCommand(t *testing.T) {
	_, ok := IsOnCommand("!cmd", []string{"cmd"})
	if !ok {
		t.Error("IsOnCommand basic invalid")
	}

	_, ok = IsOnCommand("!cmd", []string{"cmdcmd"})
	if ok {
		t.Error("IsOnCommand including cmd fail")
	}

	_, ok = IsOnCommand("!chicken", []string{"aeiouy"})
	if ok {
		t.Error("IsOnCommand invalid cmd fail")
	}
}

func TestIsOnCommandParams(t *testing.T) {
	params, ok := IsOnCommand("!cmd param1 param2", []string{"cmd"})
	if !ok {
		t.Error("IsOnCommand with params failed parsing cmd")
	}

	if params != "param1 param2" {
		t.Error("IsOnCommand with params failed getting params")
	}
}
