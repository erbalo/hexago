package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erbalo/hexago/internal/adapters/menu"
	"github.com/erbalo/hexago/internal/adapters/repository/file/card"
	cardDomain "github.com/erbalo/hexago/internal/app/card"
)

func main() {
	cardRepo := card.NewRepository("cards.json")
	cardService := cardDomain.NewService(cardRepo)

	wizard := menu.InitialModel(*cardService)
	program := tea.NewProgram(wizard)

	if _, err := program.Run(); err != nil {
		fmt.Printf("😖, there's been an error: %v", err)
	}
}
