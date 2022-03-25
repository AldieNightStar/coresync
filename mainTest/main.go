package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/AldieNightStar/coresync"
)

func main() {
	reg := make(coresync.CommandRegistry)
	reg["abc"] = func(auth string, args []string) (int, string) {
		return len(args), strings.Join(args, "; ")
	}
	go func() {
		time.Sleep(time.Second)
		cl := coresync.NewHttpClient("http://localhost:8080")
		code, message := cl.DoRequest("abc", []string{"A", "B"})
		fmt.Println(code, message)
	}()
	coresync.ServeHttp("0.0.0.0", 8080, reg)
}
