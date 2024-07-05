package etcdx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewClients,
)
