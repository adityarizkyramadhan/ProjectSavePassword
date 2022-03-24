package service

import (
	"ProjectSavePassword/config"
	"ProjectSavePassword/model"
	"ProjectSavePassword/response"
)

var db, _ = config.InitDB()

func RegisterUser(Username string, Password string) (response.RegisterReturn, error) {
	var data response.RegisterReturn
	body := model.DataUser{Username: Username, Password: Password}
	if err := db.Create(&body).Error; err != nil {
		return data, err
	}
	data = response.RegisterReturn{
		ID:       body.ID,
		Username: Username,
		Token: ,
	}
}

// func AddPassword(UserId uint, Username string, Password string, Platform string) error {

// }
