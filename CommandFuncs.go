package coresync

type CommandFunc func(auth string, args []string) (int, string)
