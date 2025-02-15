package main

import (
	"time"

	"github.com/MechamJonathan/lotr-companion-app/internal/theoneapi"
)

func main() {
	theoneClient := theoneapi.NewClient(5 * time.Second)
	cfg := &config{
		theoneapiClient: theoneClient,
	}

	startRepl(cfg)

}
