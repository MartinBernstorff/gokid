package shell

import (
	"fmt"
	"os"
	"os/exec"
)

type Shell interface {
	Run(cmd string) error
}

type RealShell struct{}

func New() Shell {
	return &RealShell{}
}

func (s *RealShell) Run(cmd string) error {
	// Figure out the calling shell
	fmt.Print(cmd)
	shell := os.Getenv("SHELL")
	command := exec.Command(shell, "-c", cmd)

	// Set up pipes for standard input, output, and error
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Execute the command
	err := command.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
