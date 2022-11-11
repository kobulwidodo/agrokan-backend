package user

import (
	"agrokan-backend/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(user entity.User) (entity.User, error)
	Get(params entity.SelectUserParam) (entity.User, error)
}

type user struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	u := &user{
		db: db,
	}

	return u
}

func (u *user) Create(user entity.User) (entity.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&user).Error; err != nil {
		return user, err
	}

	if err := tx.Commit().Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Get(params entity.SelectUserParam) (entity.User, error) {
	user := entity.User{}

	if err := u.db.Where(&params).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
