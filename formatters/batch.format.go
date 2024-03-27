package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
	structs "github.com/we-and/weand_backend_common/structs"
	timezone "github.com/we-and/weand_backend_common/timezone"
)

type BatchFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	//	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
	//	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	Pid  uint32             `json:"pid"`
	Jobs []JobFormattedItem `json:"jobs,omitempty"`
}

func FormatBatches(items []m.SendBatch, me structs.Me, reqTzData *timezone.TzData) []BatchFormattedItem {
	res := []BatchFormattedItem{}
	for _, v := range items {
		res = append(res, FormatBatch(v, me, reqTzData))
	}
	return res
}
func FormatBatch(item m.SendBatch, me structs.Me, reqTzData *timezone.TzData) BatchFormattedItem {

	res := BatchFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		//		UpdatedAt: item.UpdatedAt,
		//		DeletedAt: item.DeletedAt,
		Jobs: FormatJobs(item.Jobs, me, reqTzData),
	}

	return res
}
