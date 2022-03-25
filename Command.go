package coresync

type CommandRegistry map[string]CommandFunc

type CommandDTO struct {
	CommandName string   `json:"command"`
	Arguments   []string `json:"args"`
	AuthString  string   `json:"auth"`
}

type ResponseDTO struct {
	ResultCode    int    `json:"code"`
	ResultMessage string `json:"message"`
}

func (r *ResponseDTO) JSON() string {
	return toJson(r)
}

func (c *CommandDTO) JSON() string {
	return toJson(c)
}

func NewCommand(auth string, name string, args []string) *CommandDTO {
	return &CommandDTO{name, args, auth}
}

func NewResponse(code int, message string) *ResponseDTO {
	return &ResponseDTO{code, message}
}
