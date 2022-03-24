package main

import (
	"ProjectSavePassword/config"
	"ProjectSavePassword/endpoint"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	endpoint.StartEndpoint()
}
