package main

import (
	"ProjectSavePassword/config"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}
}
