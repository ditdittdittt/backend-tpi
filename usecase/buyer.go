package usecase

import (
	"time"

	"github.com/palantir/stacktrace"

	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/repository"
)

type BuyerUsecase interface {
	Create(buyer *entities.Buyer) error
	Delete(id int) error
	Update(buyer *entities.Buyer) error
	GetByID(id int) (entities.Buyer, error)
	Index() (buyers []entities.Buyer, err error)
}

type buyerUsecase struct {
	BuyerRepository repository.BuyerRepository
}

func (b *buyerUsecase) Delete(id int) error {
	err := b.BuyerRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Buyer repository error")
	}

	return nil
}

func (b *buyerUsecase) Update(buyer *entities.Buyer) error {
	buyer.UpdatedAt = time.Now()

	err := b.BuyerRepository.Update(buyer)
	if err != nil {
		return stacktrace.Propagate(err, "[Update] Buyer repository error")
	}

	return nil
}

func (b *buyerUsecase) GetByID(id int) (entities.Buyer, error) {
	buyer, err := b.BuyerRepository.GetByID(id)
	if err != nil {
		return buyer, stacktrace.Propagate(err, "[GetByID] Buyer repository error")
	}

	return buyer, nil
}

func (b *buyerUsecase) Index() (buyers []entities.Buyer, err error) {
	selectedField := []string{"id", "nik", "name", "status", "address", "phone_number", "created_at", "updated_at"}

	buyers, err = b.BuyerRepository.GetWithSelectedField(selectedField)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[GetSelectedField] Buyer repository error")
	}

	return buyers, nil
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
