package server

import (
	"errors"

	"github.com/prateekpandey14/electron/electron"
	context "golang.org/x/net/context"
)

func (s ElectronServer) Read(ctx context.Context, r *electron.Request) (*electron.Response, error) {
	return nil, errors.New("not yet implemented")
}
