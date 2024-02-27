package config

import (
	"context"
	"flag"
)

func readArgs(c *context.Context) {
	var config = flag.String("config", "", "path to config")

	flag.Parse()

	if *config != "" {
		*c = context.WithValue(*c, "config", *config)
	}
}
