package coresync

import (
	"fmt"
	"log"
	"net/http"
)

// Create server
//   ip - (Could be empty "") IP which should be served (Recommended 0.0.0.0)
//   port - listening port (Could be 8080 or 8000)
//   commands - Map of [name]CommandFunc with commands in it.
//      command returns status code and string message
func Serve(ip string, port int, commmands map[string]CommandFunc) {
	http.HandleFunc("/", serverApiCreate(commmands))
	portString := fmt.Sprintf("%s:%d", ip, port)
	log.Fatal(http.ListenAndServe(portString, nil))
}

func serverApiCreate(commands map[string]CommandFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		cmd, err := readCommandDTO(r.Body)
		if err != nil {
			writeResponseString(w, http.StatusBadRequest, "Wrong JSON")
			return
		}

		fn, ok := commands[cmd.CommandName]
		if !ok {
			writeResponseString(w, http.StatusNotFound, "No such command")
			return
		}

		status, message := fn(cmd.AuthString, cmd.Arguments)

		writeResponseString(w, http.StatusOK, NewResult(status, message).JSON())
	}
}
