package handler

import (
	"ProjectSavePassword/helper"
	"ProjectSavePassword/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var body user
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.ResponseAPI("Error when parsing JSON", http.StatusUnprocessableEntity, err.Error()))
	}
	data, err := service.RegisterUser(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("Error when registering user", http.StatusInternalServerError, err.Error()))
	}
	c.JSON(http.StatusCreated, helper.ResponseAPI("User Registered", http.StatusCreated, data))
}

func LoginUser(c *gin.Context) {
	var body user
	if err := c.Bind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, helper.ResponseAPI("Error when parsing JSON", http.StatusUnprocessableEntity, err.Error()))
	}
	data, err := service.LoginUser(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("Error when querry user", http.StatusInternalServerError, err.Error()))
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("User Login", http.StatusOK, data))
}
