package domain

type TicketType string

var (
	ErrTicketPriceZero = errors.New("Ticket price must be greater than 0")
)

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID string
	EventID string
	Spot *Spot
	TicketType TicketType
	Price float64
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.price <= 0 {
		return ErrTicketPriceZero
	}
}