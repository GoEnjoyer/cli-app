package ezcli

import "errors"

var definedFlags []*Flag

type Flag struct {
	name        string
	isRequired  bool
	description string
}

func NewFlag(name string, isRequired bool, description string) *Flag {
	flag := &Flag{
		name:        name,
		isRequired:  isRequired,
		description: description,
	}

	definedFlags = append(definedFlags, flag)

	return flag
}

func (flag *Flag) getValue(args []string) (flagValue string, err error) {
	isFound := false

	for i, v := range args {
		if v == flag.name && i+1 < len(args) {
			flagValue = args[i+1]
			isFound = true
		}
	}

	if !isFound && flag.isRequired {
		return "", errors.New("flag and corresponding value not found")
	}

	return flagValue, nil
}
