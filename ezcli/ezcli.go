package ezcli

import (
	"fmt"
)

func setHelpCommand() {
	SetCommand("help", "help command", nil, func(data CmdData, err error) {
		fmt.Println("########## COMMANDS ##########")
		for _, cmd := range definedCommands {
			fmt.Printf("%v - %v\n", cmd.name, cmd.description)
		}

		fmt.Println("########## FLAGS ##########")
		for _, flag := range definedFlags {
			fmt.Printf("%v - %v\n", flag.name, flag.description)
		}
	})
}

func Run(args []string) error {
	commandName := args[1]
	args = args[2:]

	setHelpCommand()

	command, err := findCommand(commandName)
	if err != nil {
		return err
	}

	command.execute(args)

	return nil
}
