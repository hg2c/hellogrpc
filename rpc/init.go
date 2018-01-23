package rpc

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))

	logger, _ = zap.NewDevelopment()
}
