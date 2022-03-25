package coresync

type ServerCallBack func(command *CommandDTO) *ResponseDTO

func (s ServerCallBack) ToClient() *Client {
	return &Client{
		sendFunc: s,
		auth:     "",
	}
}
