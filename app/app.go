package app

import (
	"fmt"

	"github.com/hectorgabucio/tftbuddy/model"
)

type App struct {
	Collector model.Collector
}

func (a *App) Initialize() {
	fmt.Println("Initializing app...")
	decks := a.Collector.CollectDecks()
	fmt.Println(decks)
}
