package config

import (
	"context"
	"fmt"
)

func Init(c context.Context) {
	readArgs(&c)

	if c.Value("config") != "" && c.Value("config") != nil {
		fmt.Println("Loading .env")
		readConfigFile(c)
	}
}
