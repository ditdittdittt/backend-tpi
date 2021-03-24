package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type TransactionItemRepository interface {
	GetReport(tpiID int, districtID int, from string, to string) ([]map[string]interface{}, error)
	Delete(id int) error
}

type transactionItemRepository struct {
	db gorm.DB
}

func (t *transactionItemRepository) Delete(id int) error {
	err := t.db.Delete(&entities.TransactionItem{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionItemRepository) GetReport(tpiID int, districtID int, from string, to string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := `SELECT ti.id, t.code, f.name AS fisher_name, b.name AS buyer_name, ft.name AS fish_name, ft.code AS fish_code, c.weight, c.weight_unit, a.price 
		FROM transaction_items AS ti
		INNER JOIN auctions AS a ON ti.auction_id = a.id
		INNER JOIN transactions AS t ON ti.transaction_id = t.id
        INNER JOIN caughts AS c ON a.caught_id = c.id
        INNER JOIN fishers AS f ON c.fisher_id = f.id
        INNER JOIN buyers AS b ON t.buyer_id = b.id
        INNER JOIN fish_types AS ft ON c.fish_type_id = ft.id`

	if tpiID != 0 {
		query = query + " WHERE t.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS tpi ON t.tpi_id = tpi.id WHERE tpi.district_id = " + strconv.Itoa(districtID)
	}

	query = query + ` AND t.created_at BETWEEN "%s" AND "%s" AND c.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to)

	err := t.db.Raw(query).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransactionItemRepository(database gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{db: database}
}
