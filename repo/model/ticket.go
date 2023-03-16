package model

type TicketTYPE int

const (
	DISCOUNT TicketTYPE = iota
	MANJIAN
)

type Ticket struct {
	Type     TicketTYPE
	DisCount float64
	Edge     float64
	Sub      float64
}
