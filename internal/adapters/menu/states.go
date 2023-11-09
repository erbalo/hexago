package menu

type InputState int

const (
	Principal InputState = iota
	PromptingBIN
	PromptingLastDigits
	PromptingNetwork
	PromptingIssuer
)
