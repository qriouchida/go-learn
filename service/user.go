package service

import (
	"github.com/jinzhu/gorm"
	"king-tool-api/model"
	"time"
)

type Users struct {
	engine *gorm.DB
}

func NewUsers(engine *gorm.DB) *Users {
	u := Users{
		engine: engine,
	}
	return &u
}

func (u *Users) Create(input *model.UserInput) (*model.Users, error) {
	now := time.Now()
	user := model.Users{
		Name:      input.Name,
		Address:   input.Address,
		CreatedAt: now,
		UpdatedAt: now,
	}
	result := u.engine.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *Users) GetOne(userID int) (*model.Users, error) {
	user := model.Users{}
	result := u.engine.Table("users").Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *Users) GetAll() ([]*model.Users, error) {
	users := []*model.Users{}
	result := u.engine.Table("users").First(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *Users) Update(input *model.UserInput, userID int) (*model.Users, error) {
	user := model.Users{
		Name:    input.Name,
		Address: input.Address,
	}
	result := u.engine.Table("users").Where("id = ?", userID).Update(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return u.GetOne(userID)
}

func (u *Users) Delete(userID int) error {
	user := model.Users{}
	result := u.engine.Table("users").Where("id = ?", userID).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
