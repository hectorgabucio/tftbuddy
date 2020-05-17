package model

type Champion struct {
	Name  string
	Stars uint32
}

type Deck struct {
	Name      string
	Champions []Champion
}

type Collector interface {
	CollectDecks() []Deck
}
