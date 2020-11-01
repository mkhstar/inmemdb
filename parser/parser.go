package parser

import (
	"fmt"
	"strings"

	"github.com/mkhstar/inmemdb/result"

	"regexp"

	"github.com/mkhstar/inmemdb/command"
	"github.com/mkhstar/inmemdb/types"
)

// Parse parses incoming inMemDB command
func Parse(rawCmd string) (command.Command, result.Resulter, error) {
	// Check for valid flags
	includeJSONFlag := strings.Contains(rawCmd, "--json")
	var outputFormat string
	if includeJSONFlag {
		outputFormat = "json"
	} else {
		outputFormat = "string"
	}

	commandToParse := rawCmd
	if includeJSONFlag {
		commandToParse = strings.Replace(rawCmd, "--json", "", 1)
	}

	re, err := regexp.Compile(`("[^"]+?"\S*|\S+)`)
	if err != nil {
		return command.Command{OutputFormat: outputFormat}, types.InvalidCommandError(fmt.Sprint("(error) Command failed to parse")), nil
	}
	rawCmdParts := re.FindAllString(commandToParse, -1)
	cmdLen := len(rawCmdParts)

	if cmdLen == 0 {
		return command.Command{OutputFormat: outputFormat}, &result.Info{Result: ""}, nil
	}
	if cmdLen == 1 && strings.ToLower(rawCmdParts[0]) == "exit" {
		return command.Command{OutputFormat: outputFormat}, nil, types.ClientExitError("Bye!")
	}

	// validateCmd
	cmd := command.Command{
		Type:         rawCmdParts[0],
		Args:         rawCmdParts[1:],
		OutputFormat: outputFormat,
	}

	resulter := validateCmd(&cmd)

	return cmd, resulter, nil
}
