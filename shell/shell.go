package shell

import (
	"fmt"
	"os"
	"os/exec"
)

type Shell interface {
	Run(cmd string) error
	RunWithOutput(cmd string) (string, error)
}

type RealShell struct{}

func New() Shell {
	return &RealShell{}
}

func (s *RealShell) Run(cmd string) error {
	// Figure out the calling shell
	fmt.Printf("%s\n", cmd)
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
func (s *RealShell) RunWithOutput(cmd string) (string, error) {
	shell := os.Getenv("SHELL")
	command := exec.Command(shell, "-c", cmd)

	// Set up pipes for standard output and error
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
