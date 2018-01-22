package main

import (
	"runtime"

	"github.com/hg2c/hellogrpc/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
