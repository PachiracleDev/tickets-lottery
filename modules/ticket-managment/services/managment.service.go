package services

import (
	"app/modules/ticket-managment/models"
)

type TicketService interface {
	GetTickets() []models.Ticket
	GetTicket(id string) models.Ticket
	CreateTicket(ticket models.Ticket) models.Ticket
	UpdateTicket(id string, ticket models.Ticket) models.Ticket
}
