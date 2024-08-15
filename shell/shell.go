package shell

import (
	"os"
	"os/exec"
)

func Run(
	command string,
	args ...string,
) error {
	// Figure out the calling shell
	shell := os.Getenv("SHELL")
	all_args := append([]string{"-c", command}, args...)
	cmd := exec.Command(shell, all_args...)

	// Set up pipes for standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
