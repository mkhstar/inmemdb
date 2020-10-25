package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mkhstar/inmemdb/db"
	"github.com/mkhstar/inmemdb/server"
)

func main() {
	quitServer := make(chan os.Signal)
	signal.Notify(quitServer, os.Interrupt, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, os.Kill)

	go server.CreateServer()
	select {
	case <-quitServer:
		db.Persist()
	}
}
