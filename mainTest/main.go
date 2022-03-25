package main

import (
	"fmt"
	"strings"

	"github.com/AldieNightStar/coresync"
)

func main() {
	reg := make(coresync.CommandRegistry)
	reg["abc"] = func(auth string, args []string) (int, string) {
		return len(args), strings.Join(args, "; ")
	}
	// coresync.ServeHttp("0.0.0.0", 8080, reg)
	c := coresync.ServeRaw(reg)
	resp := c(&coresync.CommandDTO{
		CommandName: "abc",
		Arguments:   []string{"Arg1", "Arg2"},
		AuthString:  "",
	})
	fmt.Println(resp.ResultCode, resp.ResultMessage)
}
