package server

import (
	"fmt"
	"net/http"

	"github.com/mkhstar/inmemdb/result"
)

func output(w http.ResponseWriter, resulter result.Resulter, outputFormat string) {
	if outputFormat == "json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(200)
		fmt.Fprintln(w, resulter.JSON())
	} else {
		fmt.Fprintln(w, resulter.Echo())
	}
}
