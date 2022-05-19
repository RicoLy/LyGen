package main

import (
	"LyGen/core"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	core.Start(app)
}



