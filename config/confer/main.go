package main

import (
	"fmt"
	"time"

	"github.com/jacobstr/confer"
)

var Config *confer.Config

func init() {
	setup()
}

func setup() {

	Config = confer.NewConfig()
	appenv := "" //appenv := os.Getenv("MYAPP_ENV")
	paths := []string{"config.yaml"}

	if appenv != "" {
		paths = append(paths, fmt.Sprintf("config.%s.yaml", appenv))
	}

	if err := Config.ReadPaths(paths...); err != nil {
		fmt.Println(err)
	}
}

func main() {

	go func() {
		for {
			setup()
			time.Sleep(15 * time.Second)
		}
	}()

	for {
		Config.SetDefault("app.database.hosts", "content")
		dbhost := Config.GetString("app.database.host")
		fmt.Println("CONFIG VALUE:", dbhost)

		time.Sleep(5 * time.Second)
	}
}
