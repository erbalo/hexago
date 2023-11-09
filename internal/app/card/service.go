package card

import (
	"github.com/erbalo/hexago/internal/app/domain"
	"github.com/erbalo/hexago/internal/ports"
)

type Service struct {
	repository ports.CardRepository
}

func NewService(repo ports.CardRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) GetAll() ([]domain.CardRepresentation, error) {
	return service.repository.GetAll()
}

func (service *Service) Create(request domain.CardCreateReq) (*domain.CardRepresentation, error) {
	return service.repository.Create(request)
}
