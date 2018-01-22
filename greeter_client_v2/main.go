/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"

	"github.com/hg2c/hellogrpc/greeter"

	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
	"go.uber.org/zap"
)

func main() {
	logger0, _ := zap.NewDevelopment()
	logger := log.NewFactory(logger0.With(zap.String("service", "greeter")))

	metricsFactory := xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms
	logger.Bg().Info("Using expvar as metrics backend")

	tracer := tracing.Init("greeter", metricsFactory.Namespace("greeter", nil), logger)
	c := greeter.NewClient(tracer, logger)
	c.Hello("luo")
}
