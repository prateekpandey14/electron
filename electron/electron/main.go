package main

import (
	"github.com/lileio/lile"
	"github.com/prateekpandey14/electron/electron"
	"github.com/prateekpandey14/electron/electron/electron/cmd"
	"github.com/prateekpandey14/electron/electron/server"
	"google.golang.org/grpc"
)

func main() {
	s := &server.ElectronServer{}

	lile.Name("electron")
	lile.Server(func(g *grpc.Server) {
		electron.RegisterElectronServer(g, s)
	})

	cmd.Execute()
}
