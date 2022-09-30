package file

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
    path string
}


func decodeTickets(data string) (tickets []service.Ticket, err error) {

    linesCsv := strings.Split(data, "\n")

    tickets = make([]service.Ticket, len(linesCsv))

    //  1000,Magdalen Covelle,mcovellerr@wired.com,Jamaica,23:20,563
    for idx, strline := range linesCsv {
        
	line := strings.Split(strline, ",")

	id, err := strconv.Atoi(line[0])
	if err != nil {
	    return nil, err
	}
	
	price, err := strconv.Atoi(line[5])
	if err != nil {
	    return nil, err
	}

	tickets[idx] = service.NewTicket(id, price, line[1], line[2], line[3], line[4])
    }

    return tickets, err
}


func (f *File) Read() ([]service.Ticket, error) {

    dataByte, err := os.ReadFile(f.path)
    if err != nil {
	fmt.Printf(" %s\n ", err.Error())
	return nil, nil
    }

    data := string(dataByte)
    // fmt.Printf(" reading ... \n%s\n", data)

    tickets, err := decodeTickets(data)     
    if err != nil {
	return nil, err
    }
   
    return tickets, err
}


func (f *File) Write(tickets []service.Ticket) error {

    var lineStr string
    var lines string

    // 992,Derrek Treherne,dtrehernerj@moonfruit.com,China,13:17,1092
    for _, ticket := range tickets {
	lineStr = fmt.Sprintf("%d,%s,%s,%s,%s,%d\n",
				ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	
	lines += lineStr
    }

    fmt.Printf(" ****** %s \n", lines)
    
    

    return nil
}

func (f *File) SetPath(pathFile string) {
    f.path = pathFile
}

