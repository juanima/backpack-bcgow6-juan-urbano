package service

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	// Get all the tickets
	ReadAll() []Ticket
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func NewTicket(id, price int, names, email, destination, date string) (ticket Ticket) {

	ticket = Ticket{}
	ticket.Id = id
	ticket.Price = price
	ticket.Email = email
	ticket.Destination = destination
	ticket.Date = date
	ticket.Names = names

	return ticket
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	tickets := b.Tickets

	for _, ticket := range tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}

	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	tickets := b.Tickets
	for _, ticket := range tickets {
		if ticket.Id == id {

			ticket.Price = t.Price
			ticket.Email = t.Email
			ticket.Destination = t.Destination
			ticket.Date = t.Date
			ticket.Names = t.Names

			return ticket, nil
		}
	}

	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {

	tickets := b.Tickets
	b.Tickets = append(tickets[:id-1], tickets[(id-1)+1:]...)

	return id, nil
}

func (b *bookings) ReadAll() (tickets []Ticket) {
	return b.Tickets
}
