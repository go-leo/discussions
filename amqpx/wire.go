package amqpx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewConnections,
)
