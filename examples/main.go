package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tonkpils/go-aha/aha"
)

func main() {
	username := os.Getenv("AHA_USERNAME")
	password := os.Getenv("AHA_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("AHA_USERNAME and AHA_PASSWORD env vars required")
	}

	accountName := os.Getenv("AHA_ACCOUNT")
	if accountName == "" {
		log.Fatal("AHA_CCOUNT env var is required")
	}

	client, err := aha.NewBasicAuthClient(username, password, accountName)
	if err != nil {
		log.Fatal(err)
	}

	products, err := client.Products.ListAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		fmt.Printf("%+v\n", product)
	}
}
