package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func execute(ctx context.Context) error {
	for index := 1; index <= 10; index++ {
		select {
		case <-ctx.Done():
			time.Sleep(500 * time.Millisecond)
			return ctx.Err()
		default:
			time.Sleep(2 * time.Second)
			println("Processing item", index)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		cancel()
	}()

	if err := execute(ctx); err != nil {
		println("Execution interrupted:", err.Error())
		os.Exit(1)
	}

	println("Execution completed successfully")
}
