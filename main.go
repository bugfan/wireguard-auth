package main

import (
	"github.com/bugfan/wireguard-auth/env"
	"github.com/bugfan/wireguard-auth/srv"
)

func main() {
	// run it
	srv.NewServer(env.Get("api_addr")).Run()
}
