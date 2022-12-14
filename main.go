package main

import (
	"runtime"
	"technical-test/cmd"
)

var (
	appName    = "DOT INDONESIA"
	appVersion = "v1.0.0"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
