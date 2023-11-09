package ports

import "github.com/erbalo/hexago/internal/app/domain"

type CardRepository interface {
	GetAll() ([]domain.CardRepresentation, error)
	Create(domain.CardCreateReq) (*domain.CardRepresentation, error)
}
