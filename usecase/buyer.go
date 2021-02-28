package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type BuyerUsecase interface {
	Create(buyer *entities.Buyer) error
}

type buyerUsecase struct {
	BuyerRepository repository.BuyerRepository
}

func (b *buyerUsecase) Create(buyer *entities.Buyer) error {

	buyer.CreatedAt = time.Now()
	buyer.UpdatedAt = time.Now()

	err := b.BuyerRepository.Create(buyer)
	if err != nil {
		return stacktrace.Propagate(err, "[Create] Buyer repository error")
	}

	return nil
}

func NewBuyerUsecase(buyerRepository repository.BuyerRepository) BuyerUsecase {
	return &buyerUsecase{BuyerRepository: buyerRepository}
}
