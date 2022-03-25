package coresync

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Client struct {
	sendFunc ServerCallBack
	auth     string
}

func NewHttpClient(addr string) *Client {
	return &Client{
		sendFunc: func(command *CommandDTO) *ResponseDTO {
			sb := strings.Builder{}
			err := json.NewEncoder(&sb).Encode(command)
			if err != nil {
				return NewResponse(ResponseStatusUnknownError, "Can't convert command to JSON")
			}
			// Do reqest
			servResponse, err := http.Post(addr, "application/json", strings.NewReader(sb.String()))
			if err != nil {
				return NewResponse(ResponseStatusUnknownError, "Something wrong with response")
			}
			// Get response
			response := &ResponseDTO{}
			err = json.NewDecoder(servResponse.Body).Decode(response)
			if err != nil {
				return NewResponse(ResponseStatusUnknownError, "Response got, but can't read from JSON")
			}
			return response
		},
		auth: "",
	}
}

func NewRawClient(cb ServerCallBack) *Client {
	return cb.ToClient()
}

func NewSocketClient(addr string) *Client {
	return &Client{
		sendFunc: func(command *CommandDTO) *ResponseDTO {
			sock, err := net.Dial("tcp", addr)
			if err != nil {
				return NewResponse(ResponseStatusUnknownError, "Could not establish connection with server")
			}
			jsonStr := toJson(command)
			if jsonStr == "" {
				return NewResponse(ResponseStatusUnknownError, "Can't convert command to JSON")
			}
			fmt.Fprintln(sock, jsonStr)
			resp := &ResponseDTO{}
			err = json.NewDecoder(sock).Decode(resp)
			if err != nil {
				return NewResponse(ResponseStatusUnknownError, "Response got, but can't read from JSON")
			}
			return resp
		},
		auth: "",
	}
}

func (c *Client) DoRequest(name string, args []string) (int, string) {
	resp := c.sendFunc(&CommandDTO{
		CommandName: name,
		Arguments:   args,
		AuthString:  c.auth,
	})
	if resp.ResultCode == ResponseStatusAuthSetup {
		c.auth = resp.ResultMessage
	}
	return resp.ResultCode, resp.ResultMessage
}
