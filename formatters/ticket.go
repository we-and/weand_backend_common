package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
)

type TicketFormattedItem struct {
	CreatedAt     time.Time                    `json:"created_at"`
	IsCompleted   bool                         `json:"is_completed"`
	LastMessageId int                          `json:"last_message_id"`
	LastMessage   int                          `gorm:"foreignKey:last_message",json:"last_message,omitempty"`
	Messages      []TicketmessageFormattedItem `json:"messages"`
	Pid           uint32                       `json:"pid"`
}

func FormatTickets(items []m.Ticket) []TicketFormattedItem {
	res := []TicketFormattedItem{}
	for _, v := range items {
		res = append(res, FormatTicket(v))
	}
	return res
}
func FormatTicket(v m.Ticket) TicketFormattedItem {
	res := TicketFormattedItem{
		CreatedAt: v.CreatedAt,
		Messages:  FormatTicketmessages(v.Ticketmessages),
		Pid:       publicid.Obfuscate32bit(v.ID),
	}

	return res
}
