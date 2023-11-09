package main

import (
	"fmt"

	"github.com/erbalo/hexago/internal/adapters/http"
	"github.com/erbalo/hexago/internal/adapters/repository/memory/card"
	cardDomain "github.com/erbalo/hexago/internal/app/card"
)

func main() {
	fmt.Println("Running application...")

	memoryRepository := card.NewRepository()
	cardService := cardDomain.NewService(memoryRepository)

	server := http.NewServer(*cardService)
	server.ConfigureRoutes()
	server.Run()
}
