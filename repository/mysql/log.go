package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type LogRepository interface {
	Create(log *entities.Log) error
}

type logRepository struct {
	db gorm.DB
}

func (l *logRepository) Create(log *entities.Log) error {
	err := l.db.Create(log).Error
	if err != nil {
		return err
	}

	return nil
}

func NewLogRepository(database gorm.DB) LogRepository {
	return &logRepository{db: database}
}
