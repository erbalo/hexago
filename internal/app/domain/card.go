package domain

type CardNetwork int

const (
	Unknown CardNetwork = iota
	Visa
	Amex
	Mastercard
	Discover
)

type CardRepresentation struct {
	ID         int         `json:"id"`
	Network    CardNetwork `json:"network"`
	Bin        int         `json:"bin"`
	LastDigits int         `json:"last_digits"`
	Issuer     string      `json:"issuer"`
}

type CardCreateReq struct {
	Network    CardNetwork `json:"network" binding:"required,gte=0,lte=4"`
	Bin        int         `json:"bin"`
	LastDigits int         `json:"last_digits"`
	Issuer     string      `json:"issuer"`
}
