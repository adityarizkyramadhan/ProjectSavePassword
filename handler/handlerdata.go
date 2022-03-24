package handler

import (
	"ProjectSavePassword/helper"
	"ProjectSavePassword/model"
	"ProjectSavePassword/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type inputData struct {
	Username string `json : "username" binding:"required"`
	Password string `json : "password" binding:"required"`
	Platform string `json : "platform" binding:"required"`
}

func AddData(c *gin.Context) {
	iduser := c.MustGet("UserLoggedIn").(uint)
	var dataIn inputData
	if err := c.BindJSON(&dataIn); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.ResponseAPI("Error when parsing JSON", http.StatusUnprocessableEntity, err.Error()))
		return
	}
	data := model.SaveData{
		Username: dataIn.Username,
		Password: dataIn.Password,
		Platform: dataIn.Platform,
		UserID:   iduser,
	}
	if err := service.AddData(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("Error when registering user", http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, helper.ResponseAPI("User Registered", http.StatusCreated, nil))
}

func SearchData(c *gin.Context) {
	iduser := c.MustGet("UserLoggedIn").(uint)
	data, err := service.SearchData(iduser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("Error when querry data", http.StatusInternalServerError, err.Error()))
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data Found", http.StatusOK, data))
}
