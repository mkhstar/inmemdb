package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/mkhstar/inmemdb/result"
	"github.com/mkhstar/inmemdb/types"

	"github.com/mkhstar/inmemdb/command"
)

// Parse parses incoming inMemDB command
func Parse(body io.ReadCloser) (command.Command, result.Resulter, error) {
	var commands requestBody
	var outputFormat = "json"

	var requestBody, err = ioutil.ReadAll(body)

	if err != nil {
		return command.Command{OutputFormat: outputFormat}, types.InvalidCommandError(fmt.Sprint("(error) Command failed to parse")), nil
	}

	err = json.Unmarshal(requestBody, &commands)
	if err != nil {
		return command.Command{OutputFormat: outputFormat}, types.InvalidCommandError(fmt.Sprint("(error) Command failed to parse")), nil
	}

	outputFormat = commands.OutputFormat

	if len(commands.Command) == 0 {
		return command.Command{OutputFormat: outputFormat}, &result.Info{Result: ""}, nil
	}

	// validateCmd
	cmd := command.Command{
		Type:         commands.Command[0],
		Args:         commands.Command[1:],
		OutputFormat: outputFormat,
	}

	resulter := validateCmd(&cmd)

	return cmd, resulter, nil
}
