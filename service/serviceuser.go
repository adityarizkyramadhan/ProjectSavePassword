package service

import (
	"ProjectSavePassword/auth"
	"ProjectSavePassword/config"
	"ProjectSavePassword/model"
	"ProjectSavePassword/response"

	"golang.org/x/crypto/bcrypt"
)

var db, _ = config.InitDB()

func RegisterUser(Username string, Password string) (response.RegisterReturn, error) {
	var data response.RegisterReturn
	cryptPass, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		return data, err
	}
	body := model.DataUser{Username: Username, Password: string(cryptPass)}
	if err := db.Create(&body).Error; err != nil {
		return data, err
	}
	token, err := auth.GenerateAuthToken(body.ID)
	if err != nil {
		return data, err
	}
	data = response.RegisterReturn{
		ID:       body.ID,
		Username: Username,
		Token:    token,
	}
	return data, nil
}

func LoginUser(username string, password string) (response.RegisterReturn, error) {
	var data response.RegisterReturn
	var respon model.DataUser
	if err := db.Where("username = ?", username).Take(&respon).Error; err != nil {
		return data, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(respon.Password), []byte(password))
	if err != nil {
		return data, err
	}
	token, err := auth.GenerateAuthToken(respon.ID)
	if err != nil {
		return data, err
	}
	data = response.RegisterReturn{
		ID:       respon.ID,
		Username: username,
		Token:    token,
	}
	return data, nil
}

// func AddPassword(UserId uint, Username string, Password string, Platform string) error {

// }
