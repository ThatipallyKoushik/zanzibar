package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints"
	"github.com/uber/zanzibar/runtime"
)

var cachedServer *zanzibar.Gateway

func TestMain(m *testing.M) {
	if os.Getenv("GATEWAY_RUN_CHILD_PROCESS_TEST") != "" {
		listenOnSignals()

		code := m.Run()
		os.Exit(code)
	} else {
		os.Exit(0)
	}
}

func listenOnSignals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGUSR2)

	go func() {
		_ = <-sigs

		if cachedServer != nil {
			cachedServer.Close()
		}
	}()
}

func TestStartGateway(t *testing.T) {
	testLogger := zap.New(
		zap.NewJSONEncoder(),
		zap.Output(os.Stderr),
	)

	server, err := createGateway()
	if err != nil {
		testLogger.Error(
			"Failed to CreateGateway in TestStartGateway()",
			zap.String("error", err.Error()),
		)
		// ?
		return
	}

	cachedServer = server
	err = server.Bootstrap(endpoints.Register)
	if err != nil {
		testLogger.Error(
			"Failed to Bootstrap in TestStartGateway()",
			zap.String("error", err.Error()),
		)
		// ?
		return
	}
	logAndWait(server)
}
