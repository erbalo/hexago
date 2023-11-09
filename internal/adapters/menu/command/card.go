package command

import (
	"github.com/erbalo/hexago/internal/app/card"
	pretty "github.com/inancgumus/prettyslice"
)

type CardCommand struct {
	cardService card.Service
}

func NewCardCommand(cardService card.Service) *CardCommand {
	return &CardCommand{
		cardService,
	}
}

func (options *CardCommand) GetAll() {
	cards, _ := options.cardService.GetAll()
	pretty.Show("cards", cards)
}
