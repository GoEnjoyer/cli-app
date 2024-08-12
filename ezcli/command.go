package ezcli

import "errors"

var definedCommands []*Command

type CmdData map[string]string

func (data CmdData) Get(key string) (string, error) {
	value, ok := data[key]

	if !ok {
		return "", errors.New("key does not exist")
	}

	return value, nil
}

type Command struct {
	name        string
	description string
	flags       []*Flag
	handleFunc  func(CmdData, error)
}

func SetCommand(name string, description string, flags []*Flag, handleFunc func(CmdData, error)) {
	_, err := findCommand(name)
	if err == nil {
		panic("command already exists")
	}

	command := &Command{
		name:        name,
		description: description,
		flags:       flags,
		handleFunc:  handleFunc,
	}

	definedCommands = append(definedCommands, command)
}

func (command *Command) execute(args []string) {
	var err error
	data := make(CmdData, 2)

	if command.flags == nil {
		command.handleFunc(nil, nil)
		return
	}

	for _, flag := range command.flags {
		data[flag.name], err = flag.getValue(args)
		if err != nil {
			command.handleFunc(nil, err)
			return
		}
	}

	command.handleFunc(data, nil)
}

func findCommand(commandName string) (*Command, error) {
	for _, command := range definedCommands {
		if commandName == command.name {
			return command, nil
		}
	}

	return nil, errors.New("command not found")
}
