package main

import (
	"fmt"
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
)

func main() {
	database.ConnectDb()
	var user = models.User{}
	database.DB.First(&user)
	fmt.Println(user)
}
