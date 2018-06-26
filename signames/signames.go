// Package signames converts strings to signals.
package signames

import (
	"syscall"
)

// ParseSignalName parses a string to a signal. If the string does not represent
// a valid signal, ok is false and signal is NoSignal.
func ParseSignalName(candidate string) (signal syscall.Signal, ok bool) {
	if sig, ok := signalMap[candidate]; ok {
		return sig, true
	}
	return NoSignal, false
}

// Signal strings. Names are the same as the signal constants names in package syscall.
// The values are the names as strings.
const (
	SIGABRT   = "SIGABRT"
	SIGALRM   = "SIGALRM"
	SIGBUS    = "SIGBUS"
	SIGCHLD   = "SIGCHLD"
	SIGCLD    = "SIGCLD"
	SIGCONT   = "SIGCONT"
	SIGFPE    = "SIGFPE"
	SIGHUP    = "SIGHUP"
	SIGILL    = "SIGILL"
	SIGINT    = "SIGINT"
	SIGIO     = "SIGIO"
	SIGIOT    = "SIGIOT"
	SIGKILL   = "SIGKILL"
	SIGPIPE   = "SIGPIPE"
	SIGPOLL   = "SIGPOLL"
	SIGPROF   = "SIGPROF"
	SIGPWR    = "SIGPWR"
	SIGQUIT   = "SIGQUIT"
	SIGSEGV   = "SIGSEGV"
	SIGSTKFLT = "SIGSTKFLT"
	SIGSTOP   = "SIGSTOP"
	SIGSYS    = "SIGSYS"
	SIGTERM   = "SIGTERM"
	SIGTRAP   = "SIGTRAP"
	SIGTSTP   = "SIGTSTP"
	SIGTTIN   = "SIGTTIN"
	SIGTTOU   = "SIGTTOU"
	SIGUNUSED = "SIGUNUSED"
	SIGURG    = "SIGURG"
	SIGUSR1   = "SIGUSR1"
	SIGUSR2   = "SIGUSR2"
	SIGVTALRM = "SIGVTALRM"
	SIGWINCH  = "SIGWINCH"
	SIGXCPU   = "SIGXCPU"
	SIGXFSZ   = "SIGXFSZ"
)

var signalMap = map[string]syscall.Signal{
	SIGABRT:   syscall.SIGABRT,
	SIGALRM:   syscall.SIGALRM,
	SIGBUS:    syscall.SIGBUS,
	SIGCHLD:   syscall.SIGCHLD,
	SIGCLD:    syscall.SIGCLD,
	SIGCONT:   syscall.SIGCONT,
	SIGFPE:    syscall.SIGFPE,
	SIGHUP:    syscall.SIGHUP,
	SIGILL:    syscall.SIGILL,
	SIGINT:    syscall.SIGINT,
	SIGIO:     syscall.SIGIO,
	SIGIOT:    syscall.SIGIOT,
	SIGKILL:   syscall.SIGKILL,
	SIGPIPE:   syscall.SIGPIPE,
	SIGPOLL:   syscall.SIGPOLL,
	SIGPROF:   syscall.SIGPROF,
	SIGPWR:    syscall.SIGPWR,
	SIGQUIT:   syscall.SIGQUIT,
	SIGSEGV:   syscall.SIGSEGV,
	SIGSTKFLT: syscall.SIGSTKFLT,
	SIGSTOP:   syscall.SIGSTOP,
	SIGSYS:    syscall.SIGSYS,
	SIGTERM:   syscall.SIGTERM,
	SIGTRAP:   syscall.SIGTRAP,
	SIGTSTP:   syscall.SIGTSTP,
	SIGTTIN:   syscall.SIGTTIN,
	SIGTTOU:   syscall.SIGTTOU,
	SIGUNUSED: syscall.SIGUNUSED,
	SIGURG:    syscall.SIGURG,
	SIGUSR1:   syscall.SIGUSR1,
	SIGUSR2:   syscall.SIGUSR2,
	SIGVTALRM: syscall.SIGVTALRM,
	SIGWINCH:  syscall.SIGWINCH,
	SIGXCPU:   syscall.SIGXCPU,
	SIGXFSZ:   syscall.SIGXFSZ,
}

// NoSignal is returned when a candidate string does not represent a valid signal.
const NoSignal syscall.Signal = syscall.Signal(0)
