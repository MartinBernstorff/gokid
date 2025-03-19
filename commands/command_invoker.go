package commands

import "fmt"

func Execute(commands []Command) []error {
	// Print the description of each command
	fmt.Println("Executing commands:")
	for i, command := range commands {
		fmt.Printf("  %v. %v\n", i+1, command.action.name)
		for _, assumption := range command.assumptions {
			fmt.Printf("  Assumes: %v\n", assumption.name)
		}
	}

	var errors []error
	for _, command := range commands {
		for _, assumption := range command.assumptions {
			err := assumption.callable()
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
		fmt.Println("Executing: " + command.action.name)
		err := command.action.callable()

		if err != nil {
			// Revert commands from most recently executed to
			// least recently
			for i := range completedCommands {
				index := (len(completedCommands) - 1) - i
				cmd := completedCommands[index]

				if cmd.revert.callable == nil {
					fmt.Println(cmd.action.name + " has nothing to revert, skipping")
					continue
				}

				fmt.Println("Reverting: " + cmd.action.name)
				err := completedCommands[index].revert.callable()
				if err != nil {
					fmt.Println("Error reverting: " + err.Error())
					return []error{err}
				}
			}

			return []error{err}
		}

		completedCommands = append(completedCommands, command)
	}

	fmt.Println("All commands executed successfully")

	return nil
}
