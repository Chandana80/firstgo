package entity

type IusResponse struct {
	IamTicket IamTicket
}

type IamTicket struct {
	Ticket string
	UserId string
}