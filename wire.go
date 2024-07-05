package pkg

import (
	"github.com/go-leo/pkg/actuatorx"
	"github.com/go-leo/pkg/amqpx"
	"github.com/go-leo/pkg/configx"
	"github.com/go-leo/pkg/consulx"
	"github.com/go-leo/pkg/databasex"
	"github.com/go-leo/pkg/elasticsearchx"
	"github.com/go-leo/pkg/grpcx"
	"github.com/go-leo/pkg/kafkax"
	"github.com/go-leo/pkg/mongox"
	"github.com/go-leo/pkg/nacosx"
	"github.com/go-leo/pkg/otelx"
	"github.com/go-leo/pkg/redisx"
	"github.com/go-leo/pkg/registryx"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	actuatorx.Provider,
	amqpx.Provider,
	configx.Provider,
	databasex.Provider,
	grpcx.Provider,
	kafkax.Provider,
	nacosx.Provider,
	otelx.Provider,
	redisx.Provider,
	registryx.Provider,
	mongox.Provider,
	elasticsearchx.Provider,
	consulx.Provider,
)
