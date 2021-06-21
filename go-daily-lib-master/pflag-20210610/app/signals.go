package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)
var shutdownSignals = []os.Signal{os.Interrupt}
var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

// requestShutdown emulates a received event that is considered as shutdown signal (SIGTERM/SIGINT)
// This returns whether a handler was notified
func requestShutdown() bool {
	if shutdownHandler != nil {
		select {
		case shutdownHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}

	return false
}

// withSignals returns a context that is canceled with any signal in sigs.
func withSignals(ctx context.Context, sigs ...os.Signal) context.Context {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sigs...)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			return
		}
	}()
	return ctx
}

// WithStandardSignals cancels the context on os.Interrupt, syscall.SIGTERM.
func WithStandardSignals(ctx context.Context) context.Context {
	return withSignals(ctx, os.Interrupt, syscall.SIGTERM)
}
