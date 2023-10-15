package main

import (
	"log"
	"s4m/internal/pkg/app"

	_ "github.com/mailru/go-clickhouse"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
