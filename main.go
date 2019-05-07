package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cmattoon/conntrackr/cmd"
	"github.com/cmattoon/conntrackr/config"
)

// /proc/net/nf_conntrack - List of connections
func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			sig := <-sigchan
			fmt.Println("Caught signal: ", sig)
			os.Exit(0)
		}
	}()

	cfg := config.New()

	app := cmd.NewFromConfig(cfg)
	app.Run()
}
