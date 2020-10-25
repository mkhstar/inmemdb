package types

import (
	"encoding/json"
)

// ParseError error occured when parsing command
type ParseError string

func (parseError ParseError) Error() string {
	return string(parseError)
}

// ClientExitError error occurs when client enter the 'exit' command
type ClientExitError string

func (clientExitError ClientExitError) Error() string {
	return string(clientExitError)
}

// EmptyResultError gets thrown when result set is empty. eg a key which doesn't exist
type EmptyResultError string

func (emptyResultError EmptyResultError) Error() string {
	return string(emptyResultError)
}

// InvalidCommandError represents and invalid command passed by the client
type InvalidCommandError string

func (invalidCommandError InvalidCommandError) Error() string {
	return string(invalidCommandError)
}

// Echo implement Resulter
func (invalidCommandError InvalidCommandError) Echo() string {
	return string(invalidCommandError)
}

// JSON implement Resulter
func (invalidCommandError InvalidCommandError) JSON() string {
	result := map[string]string{
		"status": "error",
		"error":  string(invalidCommandError),
	}
	respByte, _ := json.Marshal(result)
	return string(respByte)
}
