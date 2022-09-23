package main

import "github.com/carlmjohnson/versioninfo"

type ProgramArgs struct {
	Word    string `arg:"positional, required"`
	ListAll bool   `arg:"-a,--list-all" default:"false"`
}

func (ProgramArgs) Version() string {
	return versioninfo.Short()
}
