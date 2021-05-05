package main

import (
	"atlas-mdc/kafka/consumer"
	"atlas-mdc/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.CreateLogger()

	consumer.CreateEventConsumers(l)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
}
