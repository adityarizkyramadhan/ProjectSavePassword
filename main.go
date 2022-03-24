package main

import (
	"ProjectSavePassword/config"
	"ProjectSavePassword/endpoint"
	"fmt"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	endpoint.StartEndpoint()
	fmt.Println("END")
}
