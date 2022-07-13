package main

import (
	"log"
	rest "portfolio/cmd/rest_api"
	config2 "portfolio/shared/config"
)

func main() {
	// Config file
	config := new(config2.Config)
	config.Application.Port = "7770"
	config.ThirdParty.Twitter.AuthToken = ""
	config.ThirdParty.Twitter.Host = "https://api.twitter.com/2/users/2244994945/tweets"

	// Create App
	app, err := rest.New(config)
	if err != nil {
		log.Print(err)
	}

	// Run App
	if err := app.Run(); err != nil {
		log.Print(err)
	}
}
