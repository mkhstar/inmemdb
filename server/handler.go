package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mkhstar/inmemdb/db"
	"github.com/mkhstar/inmemdb/parser"
	"github.com/mkhstar/inmemdb/platform"
	"github.com/mkhstar/inmemdb/result"
	"github.com/mkhstar/inmemdb/types"
)

func handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cmd, parseResulter, err := parser.Parse(r.Body)

	if err != nil {
		if _, ok := err.(types.ClientExitError); ok {
			fmt.Fprintf(w, "%s%s> ", err, platform.LineBreak)
		}
	} else if parseResulter != nil {
		output(w, parseResulter, cmd.OutputFormat)
	} else {
		resulter := make(chan result.Resulter)
		go db.Execute(&cmd, resulter)

		output(w, <-resulter, cmd.OutputFormat)

	}

}
