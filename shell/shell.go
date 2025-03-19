package shell

import (
	"fmt"
	"os"
	"os/exec"
)

type Shell interface {
	Run(cmd string) (string, error)
	RunQuietly(cmd string) (string, error)
}

type RealShell struct{}

func New() Shell {
	return &RealShell{}
}

func (s *RealShell) RunQuietly(cmd string) (string, error) {
	return s.run(cmd, true)
}

func (s *RealShell) Run(cmd string) (string, error) {
	return s.run(cmd, false)
}

func (s *RealShell) run(cmd string, quiet bool) (string, error) {
	// Figure out the calling shell
	fmt.Printf("%s\n", cmd)
	shell := os.Getenv("SHELL")

	// Set up pipes for standard input, output, and error
	command := exec.Command(shell, "-c", cmd)

	if !quiet {
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}

	// Execute the command
	output, err := command.Output()
	return string(output), err
}
