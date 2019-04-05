// +build echo

package main

import (
	"github.com/TGSAN/go-tun2socks/core"
	"github.com/TGSAN/go-tun2socks/proxy/echo"
)

func init() {
	registerHandlerCreater("echo", func() {
		core.RegisterTCPConnHandler(echo.NewTCPHandler())
		core.RegisterUDPConnHandler(echo.NewUDPHandler())
	})
}
