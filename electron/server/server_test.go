package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/prateekpandey14/electron/electron"
)

var s = ElectronServer{}
var cli electron.ElectronClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		electron.RegisterElectronServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = electron.NewElectronClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
