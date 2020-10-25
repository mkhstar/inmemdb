package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mkhstar/inmemdb/platform"
)

// CreateServer creates inMemDB server
// by listening on the default port 9005 since it is not taken by any known app for now,
// else the port is taken from the first argument specified
func CreateServer() {
	port := ":9005"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("MEMDB server started on port %s, Happy Hacking!%s", port, platform.LineBreak)

	for {
		client, err := conn.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(client)
	}
}
