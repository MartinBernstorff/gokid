package shell

import (
	"os"
	"os/exec"
)

func Run(
	command string,
	args ...string,
) error {
	print("\nRunning ", command, args)
	shell := os.Getenv("SHELL")
	all_args := append([]string{"-c", command}, args...)
	cmd := exec.Command(shell, all_args...)

	// Migrate the path to the new command

	// Set up pipes for standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		print("\nCommand failed with error:", err)
		panic(err)
	}

	print("\nCommand finished successfully\n\n")
	return nil
}
