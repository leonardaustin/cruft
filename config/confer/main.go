package main

import (
	"fmt"

	"github.com/jacobstr/confer"
)

var Config *confer.Config

func init() {
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
	Config.SetDefault("app.database.hosts", "content")
	dbhost := Config.GetString("app.database.hosts")
	fmt.Println(dbhost)
}
