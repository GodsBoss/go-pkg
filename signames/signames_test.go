package signames_test

import (
	"github.com/GodsBoss/go-pkg/signames"

	"syscall"
	"testing"
)

func TestSuccessfulParsing(t *testing.T) {
	sig, ok := signames.ParseSignalName(signames.SIGHUP)

	if !ok {
		t.Errorf("Expected '%s' to be parsed successfully", signames.SIGHUP)
	}
	if sig != syscall.SIGHUP {
		t.Errorf("Expected parsed signal to be %s, but got %s", syscall.SIGHUP, sig)
	}
}

func TestUnsuccessfulParsing(t *testing.T) {
	candidate := "I am not a signal"
	sig, ok := signames.ParseSignalName(candidate)

	if ok {
		t.Errorf("Exepcted '%s' not to be parsed successfully", candidate)
	}
	if sig != signames.NoSignal {
		t.Errorf("Expected parsed signal to be %s, but got %s", signames.NoSignal, sig)
	}
}
