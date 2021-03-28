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
	Get(query map[string]interface{}, date string) (caughts []entities.Caught, err error)
	Search(query map[string]interface{}) (caughts []entities.Caught, err error)
	Delete(id int) error

	// Report
	GetWeightTotal(fishTypeID int, tpiID int, from string, to string) (float64, error)
	GetFisherTotal(status string, tpiID int, from string, to string) (int, error)

	// Dashboard
	GetFisherTotalDashboard(tpiID int, districtID int) ([]map[string]interface{}, error)
	GetProductionTotalDashboard(tpiID int, districtID int, queryType string, date string) (float64, error)
	GetProductionTotalGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error)
}

type caughtRepository struct {
	db gorm.DB
}

func (c *caughtRepository) GetProductionTotalGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := ` COALESCE( 
       SUM(
           CASE
          	WHEN c.weight_unit = "Ton" THEN c.weight * 1000
          	WHEN c.weight_unit = "Kwintal" THEN c.weight * 100
			WHEN c.weight_unit = "Kg" THEN c.weight * 1
    	END), 0) AS total
		FROM caughts AS c
		INNER JOIN tpis AS t ON c.tpi_id = t.id
		INNER JOIN fish_types AS ft ON c.fish_type_id = ft.id`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	switch queryType {
	case "daily":
		query = `SELECT ft.name AS name,` + query + ` AND DATE(c.created_at) = DATE("%s") AND c.caught_status_id = 3 GROUP BY (ft.name) ORDER BY total DESC LIMIT 10`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = `SELECT DAY(c.created_at) AS date,` + query + ` AND MONTH(c.created_at) = MONTH("%s") AND YEAR(c.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (DAY(c.created_at))`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = `SELECT MONTH(c.created_at) AS month,` + query + ` AND YEAR(c.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (MONTH(c.created_at))`
		query = fmt.Sprintf(query, date)
	}

	err := c.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *caughtRepository) GetProductionTotalDashboard(tpiID int, districtID int, queryType string, date string) (float64, error) {
	var result float64

	query := `SELECT COALESCE( 
    			SUM(
       				CASE
     					WHEN c.weight_unit = "Ton" THEN c.weight * 1000
     					WHEN c.weight_unit = "Kwintal" THEN c.weight * 100
       					WHEN c.weight_unit = "Kg" THEN c.weight * 1
				END), 0) AS total
                FROM caughts AS c`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON c.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	switch queryType {
	case "daily":
		query = query + ` AND DATE(c.created_at) = DATE("%s")`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = query + ` AND MONTH(c.created_at) = MONTH("%s") AND YEAR(c.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = query + ` AND YEAR(c.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date)
	}

	query = query + " AND c.caught_status_id = 3"

	err := c.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *caughtRepository) GetFisherTotalDashboard(tpiID int, districtID int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := `SELECT f.status AS status, COALESCE(COUNT(DISTINCT c.fisher_id), 0) AS total
		FROM caughts AS c
		INNER JOIN fishers AS f ON c.fisher_id = f.id`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON c.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}
	query = query + " GROUP BY f.status"

	err := c.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *caughtRepository) GetFisherTotal(status string, tpiID int, from string, to string) (int, error) {
	var result int
	query := `SELECT COALESCE(COUNT(DISTINCT c.fisher_id), 0) 
		FROM caughts AS c
		INNER JOIN caught_items AS ci ON ci.caught_id = c.id`

	switch status {
	case "Tetap":
		query = query + " INNER JOIN fishers AS f ON c.fisher_id = f.id AND f.tpi_id = " + strconv.Itoa(tpiID)
	case "Pendatang":
		query = query + " INNER JOIN fisher_tpis AS ft ON c.fisher_id = ft.fisher_id = " + strconv.Itoa(tpiID)
	}

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	query = query + ` AND c.created_at BETWEEN "%s" AND "%s" AND ci.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to)

	err := c.db.Debug().Raw(query).Scan(&result).Error
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
 				WHEN ci.weight_unit = "Ton" THEN ci.weight * 1000
 				WHEN ci.weight_unit = "Kwintal" THEN ci.weight * 100
    			WHEN ci.weight_unit = "Kg" THEN ci.weight * 1
 				END), 0) AS total
			FROM caught_items AS ci
			INNER JOIN caughts AS c ON ci.caught_id = c.id`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if fishTypeID != 0 {
		query = query + ` AND ci.fish_type_id = ` + strconv.Itoa(fishTypeID)
	}

	query = query + ` AND c.created_at BETWEEN "%s" AND "%s" AND ci.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to)

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

func (c *caughtRepository) Get(query map[string]interface{}, date string) (caughts []entities.Caught, err error) {
	err = c.db.Where("created_at BETWEEN ? AND ?", date).
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
