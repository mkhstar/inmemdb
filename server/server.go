package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
	"github.com/mkhstar/inmemdb/db"
)

// CreateServer creates inMemDB server
// by listening on the default port 9005 since it is not taken by any known app for now,
// else the port is taken from the first argument specified

func CreateServer() {
	quitServer := make(chan os.Signal)
	signal.Notify(quitServer, os.Interrupt, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, os.Kill)

	go func() {
		<-quitServer
		db.Persist()
		os.Exit(1)
	}()

	port := ":9005"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	router := httprouter.New()

	router.POST("/", handle)
	go log.Println("Server Started on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}
