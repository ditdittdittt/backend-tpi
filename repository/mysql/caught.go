package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type CaughtRepository interface {
	GetByID(id int) (caught entities.Caught, err error)
	Create(caught *entities.Caught) error
	Update(caught *entities.Caught, data map[string]interface{}) error
	BulkUpdate(id []int, data map[string]interface{}) error
	Get(query map[string]interface{}, startDate string, toDate string) (caughts []entities.Caught, err error)
	Search(query map[string]interface{}) (caughts []entities.Caught, err error)
	Delete(id int) error
	GetWeightTotal(fishTypeID int, tpiID int, from string, to string) (float64, error)
	GetFisherTotal(status string, tpiID int, from string, to string) (int, error)
}

type caughtRepository struct {
	db gorm.DB
}

func (c *caughtRepository) GetFisherTotal(status string, tpiID int, from string, to string) (int, error) {
	var result int
	query := `SELECT COALESCE(COUNT(DISTINCT c.fisher_id), 0) 
		FROM caughts AS c 
		INNER JOIN fishers AS f ON c.fisher_id = f.id
		WHERE c.created_at BETWEEN "%s" AND "%s" AND f.status = "%s" AND c.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to, status)

	if tpiID != 0 {
		query = query + " AND c.tpi_id = " + strconv.Itoa(tpiID)
	}

	err := c.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *caughtRepository) GetWeightTotal(fishTypeID int, tpiID int, from string, to string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(	
				SUM(
    			CASE
 				WHEN weight_unit = "Ton" THEN weight * 1000
 				WHEN weight_unit = "Kwintal" THEN weight * 100
    			WHEN weight_unit = "Kg" THEN weight * 1
 				END), 0) AS total
			FROM caughts WHERE created_at BETWEEN "%s" AND "%s" AND caught_status_id = 3`

	query = fmt.Sprintf(query, from, to)

	if fishTypeID != 0 {
		query = query + ` AND fish_type_id = ` + strconv.Itoa(fishTypeID)
	}

	if tpiID != 0 {
		query = query + " AND tpi_id = " + strconv.Itoa(tpiID)
	}

	err := c.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *caughtRepository) Delete(id int) error {
	err := c.db.Delete(&entities.Caught{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *caughtRepository) Search(query map[string]interface{}) (caughts []entities.Caught, err error) {
	err = c.db.Preload("FishType").Preload("Fisher").Find(&caughts, query).Error
	if err != nil {
		return nil, err
	}
	return caughts, nil
}

func (c *caughtRepository) Get(query map[string]interface{}, startDate string, toDate string) (caughts []entities.Caught, err error) {
	err = c.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Preload("Fisher").
		Preload("FishType").
		Preload("FishingGear").
		Preload("FishingArea").
		Preload("CaughtStatus").
		Find(&caughts, query).Error
	if err != nil {
		return nil, err
	}

	return caughts, nil
}

func (c *caughtRepository) BulkUpdate(id []int, data map[string]interface{}) error {
	err := c.db.Table("caughts").Where("id IN ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *caughtRepository) GetByID(id int) (caught entities.Caught, err error) {
	err = c.db.First(&caught, id).Error
	if err != nil {
		return entities.Caught{}, err
	}
	return caught, nil
}

func (c *caughtRepository) Update(caught *entities.Caught, data map[string]interface{}) error {
	err := c.db.Model(&caught).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *caughtRepository) Create(caught *entities.Caught) error {
	err := c.db.Create(&caught).Error
	return err
}

func NewCaughtRepository(database gorm.DB) CaughtRepository {
	return &caughtRepository{db: database}
}
