package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

const (
	PATH = "./tickets.csv"
)

func main() {
	var (
		tickets []service.Ticket
		data    file.File
	)

	// Funcion para obtener tickets del archivo csv
	data.SetPath(PATH)
	tickets, err := data.Read()
	if err != nil {
		panic(err.Error())
	}

	service.NewBookings(tickets)

	bookings := service.NewBookings(tickets)
	fmt.Printf("-----> %v\n", bookings)

	// 992,Derrek Treherne,dtrehernerj@moonfruit.com,China,13:17,1092
	ticket := service.NewTicket(1001, 1030, "Maria Jose", "mariajose@gmail.com", "Argentina", "14:03")
	_, err = bookings.Create(ticket)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("-----> %v\n", bookings)

	ticket, err = bookings.Read(800)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("-----> %v\n", ticket)

	ticket.Price = 100011
	ticket.Email = "juanRoa@gmail.com"

	ticket, err = bookings.Update(800, ticket)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("-----> %v\n", ticket)

	id, err := bookings.Delete(1001)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("-----> %v\n", id)
	fmt.Printf("-----> %v\n", bookings)

	// err = data.Write(bookings.GetTickets())
	// if err != nil {
	//     panic(err.Error())
	// }
}
