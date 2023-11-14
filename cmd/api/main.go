package main

import (
	cardAdapter "github.com/erbalo/hexago/internal/adapters/lambda/card"
	"github.com/erbalo/hexago/internal/adapters/repository/file/card"
	cardDomain "github.com/erbalo/hexago/internal/app/card"
)

func main() {
	cardRepo := card.NewRepository("cards.json")
	cardService := cardDomain.NewService(cardRepo)
	handler := cardAdapter.NewGetAllHandler(*cardService)
	handler.Start()
}
