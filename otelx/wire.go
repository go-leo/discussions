package otelx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewTracerProvider,
	NewMeterProvider,
)
