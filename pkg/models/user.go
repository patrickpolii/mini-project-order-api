package models

import (
	"fmt"
	"html"

	"order-api/pkg/config"
	"order-api/pkg/logger"
	"order-api/pkg/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID       uint   `json:"id" gorm:"column:id"`
	UserName string `gorm:"size:255;not null;unique" json:"user_name"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(user_name string, password string) (string, error) {

	var err error

	u := Users{}

	err = config.DB.Model(Users{}).Where("user_name = ?", user_name).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	fmt.Println("errrrr{}", err)

	if err != nil {
		logger.Error("Error Bcrypt Password " + err.Error())
		return "po", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		logger.Error("error generate token " + err.Error())
		return "", err
	} else {
		return token, nil
	}

}

func (u *Users) SaveUser() (*Users, error) {

	var err error
	errB := u.BeforeSave()
	if errB != nil {
		return u, errB
	}
	err = config.DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}

func (u *Users) BeforeSave() error {

	//mengubah password menjadi hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("errorHashingPasword " + err.Error())
		return err
	}
	u.Password = string(hashedPassword)

	//menghilangkan space pada username
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))

	return nil

}
