package shell

import (
	"bytes"
	"fmt"
	"io"
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
	shellEnv := os.Getenv("SHELL")

	var shell string
	switch shellEnv {
	case "":
		shell = "/bin/sh"
	default:
		shell = shellEnv
	}

	command := exec.Command(shell, "-c", cmd)

	// Create a buffer to store the output
	var buf bytes.Buffer

	// Create a multi-writer to write to both the terminal and the buffer
	if !quiet {
		command.Stdout = io.MultiWriter(os.Stdout, &buf)
		command.Stderr = io.MultiWriter(os.Stderr, &buf)
	} else {
		command.Stdout = &buf
		command.Stderr = &buf
	}

	// Execute the command
	err := command.Run()
	if err != nil {
		return buf.String(), fmt.Errorf("%s", buf.String())
	}
	return buf.String(), err
}
