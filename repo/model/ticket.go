package model

type TicketTYPE int

const (
	DISCOUNT TicketTYPE = iota
)

type Ticket struct {
	Type TicketTYPE
}
