package main

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
)

func main() {
	database.ConnectDb()
	InitRouter()
}
