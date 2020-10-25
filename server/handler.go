package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/mkhstar/inmemdb/platform"

	"github.com/mkhstar/inmemdb/db"
	"github.com/mkhstar/inmemdb/parser"
	"github.com/mkhstar/inmemdb/types"
)

func handle(client net.Conn) {
	defer client.Close()
	fmt.Fprintf(client, "Dear Client, Welcome to IN-MEMDB%s> ", platform.LineBreak)

	scanner := bufio.NewScanner(client)

L:
	for scanner.Scan() {

		cmd, parseResulter, err := parser.Parse(scanner.Text())

		if err != nil {
			if _, ok := err.(types.ClientExitError); ok {
				fmt.Fprintf(client, "%s%s> ", err, platform.LineBreak)
				break L
			}
		} else if parseResulter != nil {
			output(client, parseResulter, cmd.OutputFormat)
		} else {
			resulter := db.Execute(&cmd)

			output(client, resulter, cmd.OutputFormat)

		}

	}
}
