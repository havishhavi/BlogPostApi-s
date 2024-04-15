package main

import (
	"os"

	"www.blog.com/config"
	"www.blog.com/pkg/cmd"
	"www.blog.com/router"
)

func main() {
	//r := gin.Default()
	r := router.SetUpRouter()
	//why args and what is this?
	//through cli we migrate the database model to sql
	args := os.Args

	if len(args) > 1 {
		cmd.Execute()
		os.Exit(1)
	}

	r.Run(config.GetEnvWithKey("APP_Domain", "localhost") + ":" + config.GetEnvWithKey("APP_PORT", "8085"))
}
