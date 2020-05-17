package model

type Champion struct {
	Name  string
	Stars int
}

type Deck struct {
	Name      string
	Champions []Champion
}

type Collector interface {
	CollectDecks() []Deck
}
