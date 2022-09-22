package main

var args struct {
	WordToFind string `arg:"positional, required"`
	ListAll    bool   `arg:"-a,--list-all" default:"false"`
}
