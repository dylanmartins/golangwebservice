package main

import (
	"fmt"

	"github.com/dylanmartins/golangwebservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Dylan",
		LastName:  "Andrade",
	}
	fmt.Println(u)
}
