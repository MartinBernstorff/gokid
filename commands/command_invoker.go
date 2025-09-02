package commands

import (
	"fmt"
	"time"
)

func Execute(commands []Command) []error {
	// Print the description of each command
	fmt.Println("Plan:")
	for i, command := range commands {
		fmt.Printf("  %v. %v\n", i+1, command.Action.Name)
		for _, assumption := range command.Assumptions {
			fmt.Printf("    Assumes: %v\n", assumption.Name)
		}
	}
	fmt.Println("")

	var errors []error
	for _, command := range commands {
		for _, assumption := range command.Assumptions {
			err := assumption.Callable()
			if err != nil {
				errors = append(errors, err)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	var completedCommands []Command
	for _, command := range commands {
		start := time.Now()
		err := command.Action.Callable()
		duration := time.Since(start)

		if err != nil {
			fmt.Println("--- Error executing: " + command.Action.Name + " ---")
			fmt.Printf("Error: %v", err)
			fmt.Println("!!! Reverting")
			// Revert commands from most recently executed to
			// least recently
			for i := range completedCommands {
				index := (len(completedCommands) - 1) - i
				cmd := completedCommands[index]

				fmt.Println("Reverting: " + cmd.Action.Name)
				if cmd.Revert.Callable == nil {
					fmt.Println(cmd.Action.Name + " has nothing to revert, skipping")
					continue
				}

				err := completedCommands[index].Revert.Callable()
				if err != nil {
					fmt.Println("––– Error reverting: " + err.Error() + " –––")
					return []error{err}
				}
			}

			fmt.Println("--- Reverted succesfully –––")
			return []error{err}
		}

		fmt.Printf("Completed: %s (took %v)\n", command.Action.Name, duration)

		completedCommands = append(completedCommands, command)
	}

	return nil
}
