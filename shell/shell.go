package shell

import (
	"os"
	"os/exec"
)

func Run(
	command string,
) error {
	// Figure out the calling shell
	cmd := exec.Command(os.Getenv("SHELL") + " -c \"" + command + "\"")

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
