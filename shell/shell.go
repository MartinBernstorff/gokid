package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(
	command string,
	args ...string,
) error {
	cmd := exec.Command(command, args...)

	// Inherit the current environment
	cmd.Env = os.Environ()

	// Inherit the current working directory
	cmd.Dir, _ = os.Getwd()

	// Set up pipes for standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}
