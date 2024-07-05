package kafkax

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewConsumers,
	NewProducers,
)
