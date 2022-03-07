package main

import (
	"fmt"

	"github.com/octavio-luna/api_campo/db"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
}
