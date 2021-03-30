package usecase

import (
	"time"

	"github.com/palantir/stacktrace"
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/entities"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
)

type BuyerUsecase interface {
	Create(buyer *entities.Buyer, tpiID int, status string) error
	Delete(id int) error
	Update(buyer *entities.Buyer, status string) error
	GetByID(id int, tpiID int) (entities.Buyer, error)
	Index(tpiID int) (buyers []entities.Buyer, err error)
}

type buyerUsecase struct {
	buyerRepository    mysql.BuyerRepository
	buyerTpiRepository mysql.BuyerTpiRepository
}

func (b *buyerUsecase) Delete(id int) error {
	err := b.buyerRepository.Delete(id)
	if err != nil {
		return stacktrace.Propagate(err, "[Delete] Buyer repository error")
	}

	return nil
}

func (b *buyerUsecase) Update(buyer *entities.Buyer, status string) error {
	// insert log
	err := helper.InsertLog(buyer.ID, constant.Buyer)
	if err != nil {
		return err
	}
	buyer.UpdatedAt = time.Now()

	updateData := map[string]interface{}{
		"user_id":      buyer.UserID,
		"nik":          buyer.Nik,
		"name":         buyer.Name,
		"address":      buyer.Address,
		"phone_number": buyer.PhoneNumber,
	}

	existingBuyer, err := b.buyerRepository.GetByID(buyer.ID)
	if err != nil {
		return stacktrace.Propagate(err, "[GetByID] Buyer repository error")
	}

	// Permanent to temporary
	if existingBuyer.TpiID == buyer.TpiID && status == constant.TemporaryStatus {
		// remove tpi_id
		updateData["tpi_id"] = nil
		err = b.buyerRepository.Update(buyer.ID, updateData)
		if err != nil {
			return stacktrace.Propagate(err, "[Update] Buyer repository error")
		}

		// insert to buyer_tpis
		buyerTpi := &entities.BuyerTpi{
			BuyerID: buyer.ID,
			TpiID:   buyer.TpiID,
		}
		err = b.buyerTpiRepository.Create(buyerTpi)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Buyer tpi repository error")
		}
	}

	// Temporary to permanent
	if existingBuyer.TpiID != buyer.TpiID && status == constant.PermanentStatus {
		// remove buyer_tpis
		err = b.buyerTpiRepository.Delete(map[string]interface{}{"buyer_id": buyer.ID, "tpi_id": buyer.TpiID})
		if err != nil {
			return stacktrace.Propagate(err, "[Delete] Buyer tpi repository error")
		}

		// update tpi_id
		updateData["tpi_id"] = buyer.TpiID
		err = b.buyerRepository.Update(buyer.ID, updateData)
		if err != nil {
			return stacktrace.Propagate(err, "[Update] Buyer repository error")
		}
	}

	return nil
}

func (b *buyerUsecase) GetByID(id int, tpiID int) (entities.Buyer, error) {
	buyer, err := b.buyerRepository.GetByID(id)
	if err != nil {
		return buyer, stacktrace.Propagate(err, "[GetByID] Buyer repository error")
	}

	if buyer.TpiID == tpiID {
		buyer.Status = constant.PermanentStatus
	} else {
		buyer.Status = constant.TemporaryStatus
	}

	return buyer, nil
}

func (b *buyerUsecase) Index(tpiID int) (buyers []entities.Buyer, err error) {
	result := make([]entities.Buyer, 0)

	query := map[string]interface{}{
		"tpi_id": tpiID,
	}

	buyers, err = b.buyerRepository.Index(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Buyer repository error")
	}
	for _, buyer := range buyers {
		buyer.Status = "Tetap"
		result = append(result, buyer)
	}

	buyerTpis, err := b.buyerTpiRepository.Index(query)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Index] Buyer tpi repository error")
	}
	for _, buyerTpi := range buyerTpis {
		buyerTpi.Buyer.Status = "Pendatang"
		result = append(result, *buyerTpi.Buyer)
	}

	return result, nil
}

func (b *buyerUsecase) Create(buyer *entities.Buyer, tpiID int, status string) error {

	buyer.CreatedAt = time.Now()
	buyer.UpdatedAt = time.Now()

	switch status {
	case "Tetap":
		buyer.TpiID = tpiID

		err := b.buyerRepository.Create(buyer)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Buyer repository error")
		}

	case "Pendatang":
		existedBuyer, err := b.buyerRepository.Get(map[string]interface{}{"nik": buyer.Nik})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = b.buyerRepository.Create(buyer)
				if err != nil {
					return stacktrace.Propagate(err, "[Create] Buyer repository error")
				}
			} else {
				return stacktrace.Propagate(err, "[Get] Buyer repository error")
			}
		}

		if existedBuyer.Nik == buyer.Nik {
			buyer.ID = existedBuyer.ID
		}

		buyerTpi := &entities.BuyerTpi{
			BuyerID: buyer.ID,
			TpiID:   tpiID,
		}

		err = b.buyerTpiRepository.Create(buyerTpi)
		if err != nil {
			return stacktrace.Propagate(err, "[Create] Buyer tpi repository error")
		}
	}

	return nil
}

func NewBuyerUsecase(buyerRepository mysql.BuyerRepository, buyerTpiRepository mysql.BuyerTpiRepository) BuyerUsecase {
	return &buyerUsecase{buyerRepository: buyerRepository, buyerTpiRepository: buyerTpiRepository}
}
