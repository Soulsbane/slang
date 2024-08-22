package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
	Copy    bool   `arg:"-c,--copy" default:"false" help:"Copy definition to the clipboard. Does not work with --list-all"`
}

func (args ProgramArgs) Description() string {
	return "Gets the definition of a slang word from Urban Dictionary"
}

func (ProgramArgs) Version() string {
	return fmt.Sprintln("Version: ", versioninfo.Short())
}
