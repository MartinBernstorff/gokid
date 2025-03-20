package commands

import "fmt"

func Execute(commands []Command) []error {
	// Print the description of each command
	fmt.Println("Plan:")
	for i, command := range commands {
		fmt.Printf("  %v. %v\n", i+1, command.Action.Name)
		for _, assumption := range command.Assumptions {
			fmt.Printf("    Assumes: %v\n", assumption.Name)
		}
	}

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
		fmt.Println("\nExecuting: " + command.Action.Name)
		err := command.Action.Callable()

		if err != nil {
			fmt.Println("Error executing: " + err.Error())
			fmt.Println("––– Reverting –––")
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

			fmt.Println("––– Reverted succesfully –––")
			return []error{err}
		}

		completedCommands = append(completedCommands, command)
	}

	fmt.Println("All commands executed successfully")

	return nil
}
