package main

var args struct {
	WordToFind string `arg:"positional"`
	ListAll    bool   `arg:"-a,--list-all" default:"false"`
}
