package main

import (
	"context"
	"go-speedrun-control/internal/config"
	"go-speedrun-control/internal/gui"
	"os"
)

var mainContext context.Context

func main() {
	// default run context
	mainContext = context.Background()
	mainContext = context.WithValue(mainContext, "runners_count", "4")
	mainContext = context.WithValue(mainContext, "server_port", "55000")
	mainContext = context.WithValue(mainContext, "segments_count", "10")
	mainContext = context.WithValue(mainContext, "changed", "1")
	config.Init(mainContext)

	// show gui
	if os.Getenv("GUI") == "1" {
		gui.Show(&mainContext)
	}
}
