package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type AuctionRepository interface {
	Create(auction *entities.Auction) error
	GetByID(id int) (auction entities.Auction, err error)
	Search(query map[string]interface{}) (auctions []entities.Auction, err error)
	Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error)
	Delete(id int) error
	Update(auction *entities.Auction) error
	GetPriceTotal(fishTypeID int, tpiID int, districtID int, from string, to string) (float64, error)
	GetTransactionSpeed(fishTypeID int, tpiID int, districtID int, from string, to string) (float64, error)

	// Dashboard
	GetTransactionTotalDashboard(tpiID int, districtID int, queryType string, date string) (float64, error)
	GetTransactionSpeedDashboard(tpiID int, districtID int, queryType string, date string) (float64, error)
	GetTransactionTotalGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error)
	GetTransactionSpeedGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error)
}

type auctionRepository struct {
	db gorm.DB
}

func (a *auctionRepository) GetTransactionSpeedGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := ` COALESCE(AVG(
		UNIX_TIMESTAMP(a.created_at)-UNIX_TIMESTAMP(c.created_at)
	), 0) / 3600 AS speed
		FROM auctions AS a
		INNER JOIN caughts AS c ON a.caught_id = c.id
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
		query = `SELECT ft.name AS name,` + query + ` AND DATE(a.created_at) = DATE("%s") AND c.caught_status_id = 3 GROUP BY (ft.name) ORDER BY speed DESC LIMIT 10`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = `SELECT DAY(a.created_at) AS date,` + query + ` AND MONTH(a.created_at) = MONTH("%s") AND YEAR(a.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (DAY(a.created_at))`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = `SELECT MONTH(a.created_at) AS month,` + query + ` AND YEAR(a.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (MONTH(a.created_at))`
		query = fmt.Sprintf(query, date)
	}

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *auctionRepository) GetTransactionTotalGraphDashboard(tpiID int, districtID int, queryType string, date string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	query := ` COALESCE(
		SUM(a.price), 0) AS total
		FROM auctions AS a
		INNER JOIN caughts AS c ON a.caught_id = c.id
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
		query = `SELECT ft.name AS name,` + query + ` AND DATE(a.created_at) = DATE("%s") AND c.caught_status_id = 3 GROUP BY (ft.name) ORDER BY total DESC LIMIT 10`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = `SELECT DAY(a.created_at) AS date,` + query + ` AND MONTH(a.created_at) = MONTH("%s") AND YEAR(a.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (DAY(a.created_at))`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = `SELECT MONTH(a.created_at) AS month,` + query + ` AND YEAR(a.created_at) = YEAR("%s") AND c.caught_status_id = 3 GROUP BY (MONTH(a.created_at))`
		query = fmt.Sprintf(query, date)
	}

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *auctionRepository) GetTransactionSpeedDashboard(tpiID int, districtID int, queryType string, date string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(AVG(
		UNIX_TIMESTAMP(a.created_at)-UNIX_TIMESTAMP(c.created_at)
	), 0) AS result
	FROM auctions AS a
	INNER JOIN caughts AS c ON a.caught_id = c.id`

	if tpiID != 0 {
		query = query + " WHERE a.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON a.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	switch queryType {
	case "daily":
		query = query + ` AND DATE(a.created_at) = DATE("%s")`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = query + ` AND MONTH(a.created_at) = MONTH("%s") AND YEAR(a.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = query + ` AND YEAR(a.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date)
	}

	query = query + " AND c.caught_status_id = 3"

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) GetTransactionTotalDashboard(tpiID int, districtID int, queryType string, date string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(
		SUM(a.price), 0)
		FROM auctions AS a
		INNER JOIN caughts AS c ON a.caught_id = c.id`

	if tpiID != 0 {
		query = query + " WHERE a.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON a.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	switch queryType {
	case "daily":
		query = query + ` AND DATE(a.created_at) = DATE("%s")`
		query = fmt.Sprintf(query, date)
	case "monthly":
		query = query + ` AND MONTH(a.created_at) = MONTH("%s") AND YEAR(a.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date, date)
	case "yearly":
		query = query + ` AND YEAR(a.created_at) = YEAR("%s")`
		query = fmt.Sprintf(query, date)
	}

	query = query + " AND c.caught_status_id = 3"

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) GetTransactionSpeed(fishTypeID int, tpiID int, districtID int, from string, to string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(AVG(
		UNIX_TIMESTAMP(a.created_at)-UNIX_TIMESTAMP(c.created_at)
	), 0) AS result
	FROM auctions AS a
	INNER JOIN caughts AS c ON a.caught_id = c.id`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON a.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	if fishTypeID != 0 {
		query = query + " AND c.fish_type_id = " + strconv.Itoa(fishTypeID)
	}

	query = query + ` AND a.created_at BETWEEN "%s" AND "%s" AND c.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to)

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) GetPriceTotal(fishTypeID int, tpiID int, districtID int, from string, to string) (float64, error) {
	var result float64
	query := `SELECT COALESCE(
		SUM(a.price), 0) 
		FROM auctions AS a
		INNER JOIN caughts AS c ON a.caught_id = c.id`

	if tpiID != 0 {
		query = query + " WHERE c.tpi_id = " + strconv.Itoa(tpiID)
	}

	if districtID != 0 {
		query = query + " INNER JOIN tpis AS t ON a.tpi_id = t.id WHERE t.district_id = " + strconv.Itoa(districtID)
	}

	if fishTypeID != 0 {
		query = query + " AND c.fish_type_id = " + strconv.Itoa(fishTypeID)
	}

	query = query + ` AND a.created_at BETWEEN "%s" AND "%s" AND c.caught_status_id = 3`
	query = fmt.Sprintf(query, from, to)

	err := a.db.Raw(query).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *auctionRepository) Update(auction *entities.Auction) error {
	err := a.db.Model(&auction).Updates(auction).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *auctionRepository) Delete(id int) error {
	err := a.db.Delete(&entities.Auction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *auctionRepository) Get(query map[string]interface{}, startDate string, toDate string) (auctions []entities.Auction, err error) {
	err = a.db.Where("created_at BETWEEN ? AND ?", startDate, toDate).
		Where("caught_id IN (?)", a.db.Table("caughts").Select("id").Where(query)).
		Preload("Caught").
		Preload("Caught.Fisher").
		Preload("Caught.FishType").
		Preload("Caught.CaughtStatus").
		Find(&auctions).Error
	if err != nil {
		return nil, err
	}

	return auctions, nil
}

func (a *auctionRepository) Search(query map[string]interface{}) (auctions []entities.Auction, err error) {
	err = a.db.Where("caught_id IN (?)", a.db.Table("caughts").Select("id").Where(query)).
		Preload("Caught").
		Preload("Caught.Fisher").
		Preload("Caught.FishType").
		Find(&auctions).Error
	if err != nil {
		return nil, err
	}
	return auctions, nil
}

func (a *auctionRepository) GetByID(id int) (auction entities.Auction, err error) {
	err = a.db.Preload("Caught").First(&auction, id).Error
	if err != nil {
		return entities.Auction{}, err
	}
	return auction, nil
}

func (a *auctionRepository) Create(auction *entities.Auction) error {
	err := a.db.Create(&auction).Error
	return err
}

func NewAuctionRepository(database gorm.DB) AuctionRepository {
	return &auctionRepository{db: database}
}
