package main

import (
	"fmt"

	"github.com/Hassaniiii/gopher/models"
)

func main() {
	user := models.User {
		ID		: 2,
		FirstName	: "Hassan",
		LastName	: "Shahbazi",
	}

	fmt.Println(user)
}
