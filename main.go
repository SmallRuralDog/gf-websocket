package main

import (
	_ "gf-websocket/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-websocket/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
