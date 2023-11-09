package menu

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erbalo/hexago/internal/adapters/menu/command"
	"github.com/erbalo/hexago/internal/app/card"
	"github.com/erbalo/hexago/internal/app/domain"
)

type InputState int

const (
	Principal InputState = iota
	PromptingBIN
	PromptingLastDigits
	PromptingNetwork
	PromptingIssuer
)

type model struct {
	choices     []string
	cursor      int
	selected    string
	inputState  InputState
	inputs      map[string]string // Use a map to store the input fields
	cardCommand command.CardCommand
}

func InitialModel(cardService card.Service) model {
	cardCommand := command.NewCardCommand(cardService)

	return model{
		choices:     []string{"Get all cards", "Create a card"},
		cursor:      0,
		cardCommand: *cardCommand,
		inputs:      make(map[string]string),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (model model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch message := msg.(type) {
	case tea.KeyMsg:
		if model.inputState == Principal {
			switch message.String() {
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
					// Transition to the input state and begin the input sequence
					model.inputState = PromptingBIN
					return model, model.nextPromptScreen()
				}

				return model, tea.Quit
			case "q", "ctrl+c":
				return model, tea.Quit
			}
		} else {
			switch model.inputState {
			case PromptingBIN:
				if message.Type == tea.KeyRunes {
					model.inputs["bin"] += string(message.Runes)
				} else if message.Type == tea.KeyEnter {
					// Advance to the next prompt
					model.inputState = PromptingLastDigits
					return model, model.nextPromptScreen()
				} else if message.Type == tea.KeyBackspace || message.Type == tea.KeyDelete {
					// Handle backspace/delete key
					if len(model.inputs["bin"]) > 0 {
						// Remove the last character from the input
						model.inputs["bin"] = model.inputs["bin"][:len(model.inputs["bin"])-1]
					}
				}
			case PromptingLastDigits:
				if message.Type == tea.KeyRunes {
					model.inputs["last_digits"] += string(message.Runes)
				} else if message.Type == tea.KeyEnter {
					// Advance to the next prompt
					model.inputState = PromptingNetwork
					return model, model.nextPromptScreen()
				} else if message.Type == tea.KeyBackspace || message.Type == tea.KeyDelete {
					// Handle backspace/delete key
					if len(model.inputs["last_digits"]) > 0 {
						// Remove the last character from the input
						model.inputs["last_digits"] = model.inputs["last_digits"][:len(model.inputs["last_digits"])-1]
					}
				}
			case PromptingNetwork:
				if message.Type == tea.KeyRunes {
					model.inputs["network"] += string(message.Runes)
				} else if message.Type == tea.KeyEnter {
					// Advance to the next prompt
					model.inputState = PromptingIssuer
					return model, model.nextPromptScreen()
				} else if message.Type == tea.KeyBackspace || message.Type == tea.KeyDelete {
					// Handle backspace/delete key
					if len(model.inputs["network"]) > 0 {
						// Remove the last character from the input
						model.inputs["network"] = model.inputs["network"][:len(model.inputs["network"])-1]
					}
				}
			case PromptingIssuer:
				if message.Type == tea.KeyRunes {
					model.inputs["issuer"] += string(message.Runes)
				} else if message.Type == tea.KeyEnter {
					// Advance to the next prompt
					model.inputState = Principal

					bin, _ := strconv.Atoi(model.inputs["bin"])
					lastDigits, _ := strconv.Atoi(model.inputs["last_digits"])
					network, _ := strconv.Atoi(model.inputs["network"])
					issuer := model.inputs["issuer"]

					cardReq := domain.CardCreateReq{
						Bin:        bin,
						LastDigits: lastDigits,
						Network:    domain.CardNetwork(network),
						Issuer:     issuer,
					}

					model.cardCommand.Create(cardReq)

					return model, tea.Quit
				} else if message.Type == tea.KeyBackspace || message.Type == tea.KeyDelete {
					// Handle backspace/delete key
					if len(model.inputs["issuer"]) > 0 {
						// Remove the last character from the input
						model.inputs["issuer"] = model.inputs["issuer"][:len(model.inputs["issuer"])-1]
					}
				}
			}

			if message.String() == "q" || message.String() == "ctrl+c" {
				// Exit the program when 'q' or Ctrl+C is pressed
				return model, tea.Quit
			}
		}

	}

	return model, nil
}

func (m model) nextPromptScreen() tea.Cmd {
	return tea.EnterAltScreen
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

	switch model.inputState {
	case PromptingBIN:
		return "Enter BIN (int): " + model.inputs["bin"]
	case PromptingLastDigits:
		return "Enter last digits (int): " + model.inputs["last_digits"]
	case PromptingNetwork:
		return "Enter network (int): " + model.inputs["network"]
	case PromptingIssuer:
		return "Enter Issuer (int): " + model.inputs["issuer"]
	}

	return s
}
