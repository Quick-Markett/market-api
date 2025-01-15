package main

import (
	"main.go/models"
	"main.go/routers"
)

func main() {
	models.ConnectDatabase()
	r := routers.SetupRouter()
	r.Run(":8080")
}
