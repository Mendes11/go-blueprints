package main

import (
	"github.com/Mendes11/go-blueprints/simple/internal/app"
	"github.com/alecthomas/kong"
)

type cli struct {
	Debug bool           `help:"enable debug messages"`
	Cat   app.CatCommand `cmd:"" help:"Read a file"`
}

func main() {
	cli := cli{}
	ctx := kong.Parse(&cli)
	err := ctx.Run(app.CommandContext{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
