package hub

import (
	"github.com/anliksim/cmd-wrapper/internal/cmd"
	"os/exec"
)

// Command structure to wrap hub executions
type Cmd struct {
	WorkDir string ""
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
	cmd.RunCommandWithStdout(hub(e.WorkDir, "status"), e.verbose)
}

// Runs 'hub status -s'
func (e *Cmd) ShortStatus() {
	cmd.RunCommandWithStdout(hub(e.WorkDir, "status", "-s"), e.verbose)
}

// Runs 'hub checkout -B' for the provided branch name
func (e *Cmd) SwitchBranch(name string) {
	cmd.RunCommandWithStdout(hub(e.WorkDir, "checkout", "-B", name), e.verbose)
}

func hub(dir string, command ...string) *exec.Cmd {
	hubCmd := exec.Command("hub", command...)
	hubCmd.Dir = dir
	return hubCmd
}
