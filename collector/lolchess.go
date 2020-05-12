package collector

import (
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

	var decks []model.Deck

	// Find and visit all links
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		var deck model.Deck
		currentDeckName := ""
		e.ForEach(".deck-name .header-name, .tft-champion img, .avgrate span", func(_ int, elDecks *colly.HTMLElement) {
			if elDecks.Name == "td" {
				currentDeckName = strings.TrimSpace(elDecks.Text)
			} else if elDecks.Name == "img" {
				deck.Name = currentDeckName
				deck.Champions = append(deck.Champions, model.Champion{Name: elDecks.Attr("alt"), Stars: 1})
			} else {
				decks = append(decks, deck)
				deck = model.Deck{}
			}

		})
	})

	c.Visit("https://lolchess.gg/statistics/meta")

	return decks
}
