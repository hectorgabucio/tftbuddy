package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/hectorgabucio/tftbuddy/model"
)

type LolChessCollector struct {
}

func (l *LolChessCollector) CollectDecks() []model.Deck {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach(".deck-name .header-name", func(_ int, elDecks *colly.HTMLElement) {
			deckName := strings.TrimSpace(elDecks.Text)

			e.ForEach("tr .units-list", func(_ int, elChamps *colly.HTMLElement) {
				fmt.Println(deckName, elChamps.Name)
			})
		})

		/*
			e.ForEach("tr .units-list",func(_ int, el *colly.HTMLElement) {
				fmt.Println(el.Name)
			})
		*/

	})

	c.Visit("https://lolchess.gg/statistics/meta")

	return nil
}
