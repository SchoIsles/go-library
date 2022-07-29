//+build: main
package main

import (
	"context"
	"github.com/SchoIsles/go-library/pkg/logging"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logging.NewLogger()
	logger.Info("1")
	read(logger)

	ctx := context.Background()
	dict := logrus.Fields{
		"trace_id": "1232131",
	}
	ctx = logging.NewContext(ctx, dict)

	read2(ctx)
}

func read(logger logrus.FieldLogger) {
	logger.WithError(errors.WithStack(errors.New("errors"))).Error("find error")
}

func read2(ctx context.Context) {
	log := logging.DefaultLogger().WithContext(ctx)
	log.WithField("name", "abc").Info("test add field from context")
}
