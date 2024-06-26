package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
)

type SessionFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	Success   bool       `json:"success"`
}

func FormatSessions(items []m.Session) []SessionFormattedItem {
	res := []SessionFormattedItem{}
	for _, v := range items {
		res = append(res, SessionFormattedItem{
			CreatedAt: v.CreatedAt,
			Success:   v.Success,
		})
	}
	return res
}
