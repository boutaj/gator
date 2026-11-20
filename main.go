package main

import (
	"fmt"
	"log"

	"github.com/boutaj/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := config.SetUser("mouad"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}