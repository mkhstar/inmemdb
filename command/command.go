package command

import (
	"strings"
)

// Command received from client
type Command struct {
	Type         string
	Args         []string
	OutputFormat string
}

// GetType Convert command to uppercase to get key easily
func (cmd *Command) GetType() string {
	return strings.ToUpper(cmd.Type)
}
