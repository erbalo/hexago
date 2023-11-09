package card

import (
	"github.com/erbalo/hexago/internal/app/domain"
	"github.com/erbalo/hexago/internal/ports"
)

type MemoryRepository struct {
	cards []domain.CardRepresentation
}

func NewRepository() ports.CardRepository {
	return &MemoryRepository{
		cards: []domain.CardRepresentation{},
	}
}

func (repository *MemoryRepository) GetAll() ([]domain.CardRepresentation, error) {
	return repository.cards, nil
}

func (repository *MemoryRepository) Create(cardReq domain.CardCreateReq) (*domain.CardRepresentation, error) {
	var lastCard domain.CardRepresentation
	if len(repository.cards) > 0 {
		lastCard = repository.cards[len(repository.cards)-1]
	}

	var card domain.CardRepresentation
	card.ID = lastCard.ID + 1
	card.Bin = cardReq.Bin
	card.Issuer = cardReq.Issuer
	card.Network = cardReq.Network
	card.LastDigits = cardReq.LastDigits

	repository.cards = append(repository.cards, card)

	return &card, nil
}
