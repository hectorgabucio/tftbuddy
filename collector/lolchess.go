package collector

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	c.OnHTML("tbody", func(e *colly.HTMLElement) {

		trEl := e.DOM.ChildrenFiltered("tr")
		var deckName string
		trEl.Each(func(i int, s *goquery.Selection) {

			if s.AttrOr("class", "") == "deck-name" {
				deckName = strings.TrimSpace(s.Children().Filter(".header-name").Text())
			} else {
				deck := model.Deck{Name: deckName}
				unitsEl := s.ChildrenFiltered("td.units-list").ChildrenFiltered(".units").ChildrenFiltered(".unit")

				unitsEl.Each(func(j int, unit *goquery.Selection) {
					champName := unit.ChildrenFiltered(".tft-champion").ChildrenFiltered("img").AttrOr("alt", "ERROR")
					deck.Champions = append(deck.Champions, model.Champion{Name: champName, Stars: 1})
				})
				decks = append(decks, deck)
			}
		})

	})

	c.Visit("https://lolchess.gg/statistics/meta")

	return decks
}
