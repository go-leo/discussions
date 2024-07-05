package elasticsearchx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewClients,
)
