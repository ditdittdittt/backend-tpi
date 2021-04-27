package mysql

import (
	"gorm.io/gorm"

	"github.com/ditdittdittt/backend-tpi/entities"
)

type UserRepository interface {
	Create(user *entities.User) (err error)
	Update(user *entities.User) (err error)
	GetByID(id int) (user entities.User, err error)
	GetByNik(nik string) (user entities.User, err error)
	GetByUsername(username string) (user entities.User, err error)
	GetWithSelectedField(selectedField []string) (users []entities.User, err error)
	Get() (users []entities.User, err error)
}

type userRepository struct {
	db gorm.DB
}

func (u *userRepository) Get() (users []entities.User, err error) {
	err = u.db.Preload("Role").Preload("UserStatus").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetWithSelectedField(selectedField []string) (users []entities.User, err error) {
	err = u.db.Select(selectedField).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (u *userRepository) Update(user *entities.User) (err error) {
	err = u.db.Model(&user).Updates(user).Error
	return err
}

func (u *userRepository) GetByUsername(username string) (user entities.User, err error) {
	err = u.db.Where(&entities.User{
		Username: username,
	}).First(&user).Error

	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *userRepository) Create(user *entities.User) (err error) {
	err = u.db.Create(&user).Error
	return err
}

func (u *userRepository) GetByID(id int) (user entities.User, err error) {
	err = u.db.Preload("Role").Preload("Role.Permission").Preload("UserStatus").Find(&user, id).Error
	return user, err
}

func (u *userRepository) GetByNik(nik string) (user entities.User, err error) {
	err = u.db.Where(&entities.User{
		Nik: nik,
	}).Find(&user).Error
	return user, err
}

func NewUserRepository(database gorm.DB) UserRepository {
	return &userRepository{db: database}
}
