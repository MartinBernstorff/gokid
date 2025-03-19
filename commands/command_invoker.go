package commands

func Execute(commands []Command) []error {
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
			for _, completedCommand := range completedCommands {
				completedCommand.inverse()
			}
			return []error{err}
		}
		completedCommands = append(completedCommands, command)
	}

	return nil
}
