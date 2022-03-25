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

type Auth struct {
	AuthString string
}

func NewAuth(code string) *Auth {
	return &Auth{code}
}

func (a *Auth) Command(name string, args []string) *CommandDTO {
	return NewCommand(a.AuthString, name, args)
}

func NewCommand(auth string, name string, args []string) *CommandDTO {
	return &CommandDTO{name, args, auth}
}

func NewResult(code int, message string) *ResponseDTO {
	return &ResponseDTO{code, message}
}
