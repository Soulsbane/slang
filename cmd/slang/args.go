package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
}

func (args ProgramArgs) Description() string {
	return "Gets the definition of a slang word from Urban Dictionary"
}

func (ProgramArgs) Version() string {
	return fmt.Sprintln("Version: ", versioninfo.Short())
}
