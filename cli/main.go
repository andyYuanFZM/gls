// +build go1.8

package main

import (
	_ "github.com/33cn/chain33/system"
	_ "github.com/andyYuanFZM/gls/plugin"

	"github.com/33cn/chain33/util/cli"
	"github.com/andyYuanFZM/gls/cli/buildflags"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "gls")
}
