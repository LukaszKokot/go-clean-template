package main

import (
	"github.com/LukaszKokot/go-clean-template/cmd"
)

var rootCommand = cmd.RootCommand

func main() {
	rootCommand().Execute()
}
