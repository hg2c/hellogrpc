package rpc

import (
	"math/rand"
	"time"

	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

var (
	metricsBackend string
	logger         *zap.Logger
	metricsFactory metrics.Factory
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))

	metricsBackend = "expvar"
	logger, _ = zap.NewDevelopment()
	initMetrics(metricsBackend, logger)
}

// initMetrics is called before the command is executed.
func initMetrics(metricsBackend string, logger *zap.Logger) {
	if metricsBackend == "expvar" {
		metricsFactory = xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms
		logger.Info("Using expvar as metrics backend")
	} else if metricsBackend == "prometheus" {
		metricsFactory = jprom.New()
		logger.Info("Using Prometheus as metrics backend")
	} else {
		logger.Fatal("unsupported metrics backend " + metricsBackend)
	}
}
