package server

import (
	"fmt"
	"net"

	"github.com/mkhstar/inmemdb/platform"
	"github.com/mkhstar/inmemdb/result"
)

func output(client net.Conn, resulter result.Resulter, outputFormat string) {
	if outputFormat == "json" {
		fmt.Fprintf(client, "%s%s", resulter.JSON(), platform.LineBreak)
	} else {
		fmt.Fprintf(client, "%s%s> ", resulter.Echo(), platform.LineBreak)
	}
}
