package main

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"whatsm/internal/cmd"

	_ "github.com/mattn/go-sqlite3"
	_ "whatsm/internal/logic"
	_ "whatsm/internal/packed"
)

func main() {
	main, err := gcmd.NewFromObject(cmd.Main{})
	if err != nil {
		panic(err)
	}
	main.Run(gctx.New())
}
