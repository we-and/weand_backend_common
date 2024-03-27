package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
)

type LoginHistoryFormattedItem struct {
	Pid       uint32     `json:"pid"`
	CreatedAt *time.Time `json:"created_at"`
	Success   bool       `json:"success"`
}

func FormatLoginHistoryList(items []m.Loginhistory) []LoginHistoryFormattedItem {
	res := []LoginHistoryFormattedItem{}
	for _, v := range items {
		res = append(res, LoginHistoryFormattedItem{
			CreatedAt: v.CreatedAt,
			Pid:       publicid.Obfuscate32bit(v.ID),
			Success:   v.Success,
		})
	}
	return res
}
