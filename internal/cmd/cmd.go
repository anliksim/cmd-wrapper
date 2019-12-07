package cmd

import (
	"bytes"
	"log"
	"os/exec"
)

func RunCommandWithStdout(cmd *exec.Cmd, verbose bool) {
	if verbose {
		runCommandWithStdout(cmd)
	} else {
		runCommand(cmd)
	}
}

func runCommandWithStdout(cmd *exec.Cmd) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n%v", out.String())
	log.Printf("Command %v finished with error: %v", cmd.Args, err)
}

func runCommand(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
