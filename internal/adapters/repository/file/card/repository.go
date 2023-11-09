package card

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/erbalo/hexago/internal/app/domain"
	"github.com/erbalo/hexago/internal/ports"
)

type FileRepository struct {
	fileName string
}

func NewRepository(fileName string) ports.CardRepository {
	return &FileRepository{
		fileName,
	}
}

func (repository *FileRepository) GetAll() ([]domain.CardRepresentation, error) {
	var cards []domain.CardRepresentation
	file, err := os.ReadFile(repository.fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If the file doesn't exist, return an empty list
			return cards, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (repository *FileRepository) Create(cardReq domain.CardCreateReq) (*domain.CardRepresentation, error) {
	cards, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	var lastCard domain.CardRepresentation
	if len(cards) > 0 {
		lastCard = cards[len(cards)-1]
	}

	var card domain.CardRepresentation
	card.ID = lastCard.ID + 1
	card.Bin = cardReq.Bin
	card.Issuer = cardReq.Issuer
	card.Network = cardReq.Network
	card.LastDigits = cardReq.LastDigits

	cards = append(cards, card)
	data, err := json.Marshal(cards)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(repository.fileName, data, 0644)
	if err != nil {
		return nil, err
	}

	return &card, nil
}
