package main

import "github.com/AldieNightStar/coresync"

func main() {
	reg := make(coresync.CommandRegistry)
	reg["abc"] = func(auth string, args []string) (int, string) {
		return 1, "OK"
	}
	coresync.Serve("0.0.0.0", 8080, reg)
}
