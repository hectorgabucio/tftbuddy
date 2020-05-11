package main

import (
	"github.com/hectorgabucio/tftbuddy/collector"
	"github.com/hectorgabucio/tftbuddy/app"
)

func main() {

	// TODO: support more deck collectors and be able to choose between them or combinate them
	collector := collector.LolChessCollector{}

	app := &app.App{
		Collector: &collector,
	}

	app.Initialize()
}
