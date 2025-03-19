package commands

import "fmt"

func Execute(commands []Command) []error {
	// Print the description of each command
	fmt.Println("Executing commands:")
	for _, command := range commands {
		fmt.Println("- " + command.description)
	}

	var errors []error
	for _, command := range commands {
		for _, assumption := range command.assumptions {
			err := assumption()
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
		err := command.action()

		if err != nil {
			// Revert commands from most recently executed to
			// least recently
			for i := range completedCommands {
				index := (len(completedCommands) - 1) - i
				fmt.Println("Reverting: " + completedCommands[index].description)

				cmd := completedCommands[index]
				if cmd.revert == nil {
					continue
				}

				err := completedCommands[index].revert()
				if err != nil {
					fmt.Println("Error reverting: " + err.Error())
					return []error{err}
				}
			}

			return []error{err}
		}

		completedCommands = append(completedCommands, command)
	}

	return nil
}
