package menu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erbalo/hexago/internal/adapters/menu/command"
	"github.com/erbalo/hexago/internal/app/card"
)

type model struct {
	choices     []string
	cursor      int
	selected    string
	cardCommand command.CardCommand
}

func InitialModel(cardService card.Service) model {
	cardCommand := command.NewCardCommand(cardService)

	return model{
		choices:     []string{"Get all cards", "Create a card"},
		cursor:      0,
		cardCommand: *cardCommand,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (model model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if model.cursor > 0 {
				model.cursor--
			}

		case "down", "j":
			if model.cursor < len(model.choices)-1 {
				model.cursor++
			}

		case "enter", " ":
			model.selected = model.choices[model.cursor]
			// Handle the selected action
			switch model.selected {
			case "Get all cards":
				model.cardCommand.GetAll()
			case "Create a card":
				// Call your cardService method here to create a new card
			}

			return model, tea.Quit

		case "q", "ctrl+c":
			return model, tea.Quit
		}
	}

	return model, nil
}

func (model model) View() string {
	var s string
	for i, choice := range model.choices {
		// Render the cursor
		cursor := " "
		if model.cursor == i {
			cursor = cursorStyle.Render(">")
		}
		// Render the choice
		s += fmt.Sprintf("%s %s\n", cursor, choiceStyle.Render(choice))
	}
	s += continueStyle.Render("\nPress q to quit...\n\n")

	return s
}
