//+build: main
package main

import (
	"github.com/SchoIsles/go-library/pkg/logging"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func main() {
	logger := logging.NewLogger()
	logger.Info("1")
	read(logger)
	if info, ok := debug.ReadBuildInfo(); ok {
		logger.Info(info.Deps[0].Path)
	}
}

func read(logger logrus.FieldLogger) {
	logger.WithError(errors.WithStack(errors.New("errors"))).Error("find error")
}
