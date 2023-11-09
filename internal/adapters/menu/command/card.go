package command

import (
	"strconv"

	"github.com/erbalo/hexago/internal/adapters/menu/input"
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

func (command *CardCommand) Create(inputs map[string]string) {
	bin, _ := strconv.Atoi(inputs[input.BinKey])
	lastDigits, _ := strconv.Atoi(inputs[input.LastDigitsKey])
	network, _ := strconv.Atoi(inputs[input.NetworkKey])
	issuer := inputs[input.IssuerKey]

	request := domain.CardCreateReq{
		Bin:        bin,
		LastDigits: lastDigits,
		Network:    domain.CardNetwork(network),
		Issuer:     issuer,
	}

	card, _ := command.cardService.Create(request)
	cards := []*domain.CardRepresentation{card}
	pretty.Show("card", cards)
}
