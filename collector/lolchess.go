package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/hectorgabucio/tftbuddy/model"
)

type LolChessCollector struct {
}

type deckIndex struct {
	name  string
	index int
}

func (l *LolChessCollector) CollectDecks() []model.Deck {
	c := colly.NewCollector()




	// Find and visit all links
	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		e.ForEach(".deck-name .header-name, .tft-champion img, .avgrate span", func(_ int, elDecks *colly.HTMLElement) {
			if elDecks.Name == "td" {
				deckName := strings.TrimSpace(elDecks.Text)
				fmt.Println(deckName)
				//decks = append(decks, deckName)
			} else if elDecks.Name == "img" {
				fmt.Println(elDecks.Attr("alt"))
			} else {
				fmt.Print("\n\n\n")
			}

		})

	})

	c.Visit("https://lolchess.gg/statistics/meta")

	return nil
}
