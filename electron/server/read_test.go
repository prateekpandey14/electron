package server

import (
	"testing"

	"github.com/prateekpandey14/electron/electron"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestRead(t *testing.T) {
	ctx := context.Background()
	req := &electron.Request{}

	res, err := cli.Read(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
