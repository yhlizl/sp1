package main

import (
	_ "sp1/boot"
	_ "sp1/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
