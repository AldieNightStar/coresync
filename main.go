package coresync

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// Create HTTP server
//   ip - (Could be empty "") IP which should be served (Recommended 0.0.0.0)
//   port - listening port (Could be 8080 or 8000)
//   commands - Map of [name]CommandFunc with commands in it.
//      command returns status code and string message
func ServeHttp(ip string, port int, commands CommandRegistry) {
	http.HandleFunc("/", serveApiHttpCreate(commands))
	portString := fmt.Sprintf("%s:%d", ip, port)
	log.Fatal(http.ListenAndServe(portString, nil))
}

// Create Socket TCP Server
//   ip - (Could be empty "") IP which should be served (Recommended 0.0.0.0)
//   port - listening port (Could be 8080 or 8000)
//   commands - Map of [name]CommandFunc with commands in it.
//      command returns status code and string message
func ServeSocket(ip string, port int, commands CommandRegistry) error {
	portString := fmt.Sprintf("%s:%d", ip, port)
	listen, err := net.Listen("tcp", portString)
	if err != nil {
		return err
	}
	for {
		sock, err := listen.Accept()
		if err != nil {
			continue
		}
		go func() {
			buf := bufio.NewReader(sock)
			for {
				dataString, err := buf.ReadString('\n')
				if err != nil {
					break
				}
				cmd := &CommandDTO{}
				err = json.NewDecoder(strings.NewReader(dataString)).Decode(cmd)
				if err != nil {
					fmt.Fprintln(sock, NewResult(ResponseStatusUnknownError, "Bad JSON").JSON())
					continue
				}
				resp := generalApi(commands, cmd)
				fmt.Fprintln(sock, resp.JSON())
			}
		}()
	}
}

// Serve as Raw server (Without any implementation. You need to implement it by your self)
//   commands - Map of [name]CommandFunc with commands in it.
//      command returns status code and string message
func ServeRaw(commands CommandRegistry) ServerCallBack {
	return func(command *CommandDTO) *ResponseDTO {
		return generalApi(commands, command)
	}
}

func serveApiHttpCreate(commands CommandRegistry) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			writeResponseString(w, http.StatusMethodNotAllowed, "NOT ALLOWED")
			return
		}
		cmd, err := readCommandDTO(r.Body)
		if err != nil {
			writeResponseString(w, http.StatusBadRequest, "Wrong JSON: "+err.Error())
			return
		}
		resp := generalApi(commands, cmd)
		writeResponseString(w, http.StatusOK, resp.JSON())
	}
}

func generalApi(commands CommandRegistry, command *CommandDTO) *ResponseDTO {
	fn, ok := commands[command.CommandName]
	if !ok {
		return NewResult(ResponseStatusNoSuchCommand, "No such command: "+command.CommandName)
	}
	status, message := fn(command.AuthString, command.Arguments)
	return NewResult(status, message)
}
