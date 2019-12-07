package hub

import (
	"github.com/anliksim/cmd-wrapper/internal/cmd"
	"os/exec"
)

// Command structure to wrap hub executions
type Cmd struct {
	verbose bool
}

// Creates a new hub cmd wrapper
func Hub(verbose bool) *Cmd {
	return &Cmd{
		verbose: verbose,
	}
}

// Runs 'hub status'
func (e *Cmd) Status() {
	cmd.RunCommandWithStdout(hub("status"), e.verbose)
}

// Runs 'hub status -s'
func (e *Cmd) ShortStatus() {
	cmd.RunCommandWithStdout(hub("status", "-s"), e.verbose)
}

// Runs 'hub checkout -B' for the provided branch name
func (e *Cmd) SwitchBranch(name string) {
	cmd.RunCommandWithStdout(hub("checkout", "-B", name), e.verbose)
}

func hub(command ...string) *exec.Cmd {
	return exec.Command("hub", command...)
}
