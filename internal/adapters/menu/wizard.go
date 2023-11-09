package menu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erbalo/hexago/internal/adapters/menu/command"
	"github.com/erbalo/hexago/internal/adapters/menu/input"
	"github.com/erbalo/hexago/internal/app/card"
	"github.com/erbalo/hexago/internal/app/domain"
)

type model struct {
	choices     []string
	cursor      int
	selected    string
	inputState  InputState
	inputs      map[string]string
	cardCommand command.CardCommand
	createdCard domain.CardRepresentation
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputState == Principal {
			return m.updatePrincipal(msg)
		}

		if m.inputState == CreatedCard {
			card := m.createdCard
			m.cardCommand.PrintCard(card)
			return m, tea.Quit
		}

		return m.updateInput(msg)
	}

	return m, nil
}

func (m model) updatePrincipal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter", " ":
		m.selected = m.choices[m.cursor]
		return m.handleSelection()
	case "q", "ctrl+c":
		return m, tea.Quit
	}

	return m, nil
}

func (m model) handleSelection() (tea.Model, tea.Cmd) {
	switch m.selected {
	case "Get all cards":
		m.cardCommand.GetAll()
	case "Create a card":
		m.inputState = PromptingBIN
		return m, m.nextPromptScreen()
	}

	return m, tea.Quit
}

func (m model) updateInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "q" || msg.String() == "ctrl+c" {
		return m, tea.Quit
	}

	var key string
	switch m.inputState {
	case PromptingBIN:
		key = input.BinKey
	case PromptingLastDigits:
		key = input.LastDigitsKey
	case PromptingNetwork:
		key = input.NetworkKey
	case PromptingIssuer:
		key = input.IssuerKey
	default:
		return m, nil
	}

	if msg.Type == tea.KeyRunes {
		m.inputs[key] += string(msg.Runes)
	} else if msg.Type == tea.KeyEnter {
		return m.advanceInputState()
	} else if msg.Type == tea.KeyBackspace || msg.Type == tea.KeyDelete {
		if len(m.inputs[key]) > 0 {
			m.inputs[key] = m.inputs[key][:len(m.inputs[key])-1]
		}
	}

	return m, nil
}

func (m *model) advanceInputState() (tea.Model, tea.Cmd) {
	switch m.inputState {
	case PromptingBIN:
		m.inputState = PromptingLastDigits
	case PromptingLastDigits:
		m.inputState = PromptingNetwork
	case PromptingNetwork:
		m.inputState = PromptingIssuer
	case PromptingIssuer:
		m.inputState = CreatedCard
		return m.createCard()
	}

	return m, m.nextPromptScreen()
}

func (m model) createCard() (tea.Model, tea.Cmd) {
	card, _ := m.cardCommand.Create(m.inputs)
	m.createdCard = *card

	return m, tea.ExitAltScreen
}

func (m model) nextPromptScreen() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) View() string {
	var s string
	switch m.inputState {
	case Principal:
		s = m.viewPrincipal()
	case PromptingBIN:
		s = m.viewInputPrompt("Enter BIN (int): ", input.BinKey)
	case PromptingLastDigits:
		s = m.viewInputPrompt("Enter last digits (int): ", input.LastDigitsKey)
	case PromptingNetwork:
		s = m.viewInputPrompt("Enter network (int): ", input.NetworkKey)
	case PromptingIssuer:
		s = m.viewInputPrompt("Enter issuer (string): ", input.IssuerKey)
	default:
		s = ""
	}

	s += continueStyle.Render("\n\nPress q or Ctrl+c to quit...\n\n")
	return s
}

func (m model) viewPrincipal() string {
	var s string
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}
		s += fmt.Sprintf("%s %s\n", cursor, choiceStyle.Render(choice))
	}

	return s
}

func (m model) viewInputPrompt(prompt string, key string) string {
	return prompt + m.inputs[key]
}
