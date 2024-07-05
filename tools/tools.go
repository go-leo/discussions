//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/ugorji/go/codec"
	_ "google.golang.org/genproto/googleapis/rpc"
)
