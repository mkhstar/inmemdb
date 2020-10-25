package parser

import (
	"fmt"
	"strconv"

	"github.com/mkhstar/inmemdb/result"

	"github.com/mkhstar/inmemdb/command"
	"github.com/mkhstar/inmemdb/types"
)

func validateCmd(cmd *command.Command) result.Resulter {
	cmdKey := cmd.GetType()
	validCommand, ok := command.ValidCommands[cmdKey]
	if ok == false {
		return types.InvalidCommandError(fmt.Sprintf("(error) Command %s is not supported", cmdKey))
	} else if len(cmd.Args) != len(validCommand) {
		return types.InvalidCommandError(fmt.Sprintf("(error) Command %s expects only %d arguments but %d provided", cmdKey, len(validCommand), len(cmd.Args)))

	}

	for index, val := range validCommand {
		if val == "int" {
			if _, err := strconv.Atoi(cmd.Args[index]); err != nil {

				return types.InvalidCommandError(fmt.Sprintf("(error) Argument at postion %d expects an integer", index))
			}
		}
	}

	return nil
}
