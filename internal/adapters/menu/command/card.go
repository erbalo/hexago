package command

import (
	"github.com/erbalo/hexago/internal/app/card"
	"github.com/erbalo/hexago/internal/app/domain"
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

func (command *CardCommand) Create(request domain.CardCreateReq) {
	card, _ := command.cardService.Create(request)
	cards := []*domain.CardRepresentation{card}
	pretty.Show("card", cards)
}
