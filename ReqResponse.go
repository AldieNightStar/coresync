package coresync

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readCommandDTO(r io.Reader) (*CommandDTO, error) {
	cmd := &CommandDTO{}
	return cmd, json.NewDecoder(r).Decode(cmd)
}

func writeResponseString(w http.ResponseWriter, statusCode int, response string) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, response)
}
