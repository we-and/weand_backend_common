package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
	structs "github.com/we-and/weand_backend_common/structs"
	timezone "github.com/we-and/weand_backend_common/timezone"
)

type JobFormattedItem struct {
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Pid         uint32     `json:"pid"`
	Started     bool       `json:"started,omitempty"`
	Sent        bool       `json:"sent,omitempty"`
	Success     bool       `json:"success,omitempty"`
	Destination string     `json:"dest,omitempty"`
	PersonPid   uint32     `json:"person_pid,omitempty"`
	Shortmedium string     `json:"medium,omitempty"`
	Shortobject string     `json:"type,omitempty"`
}

func FormatJobs(items []m.SendJob, me structs.Me, reqTzData *timezone.TzData) []JobFormattedItem {
	res := []JobFormattedItem{}
	for _, v := range items {
		res = append(res, FormatJob(v /*,me, reqTzData*/))
	}
	return res
}
func FormatJob(item m.SendJob,

// me structs.Me,
// reqTzData *timezone.TzData
) JobFormattedItem {

	res := JobFormattedItem{
		Pid:         publicid.Obfuscate32bit(item.ID),
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		DeletedAt:   item.DeletedAt,
		Started:     item.Started,
		Sent:        item.Sent,
		Success:     item.Success,
		Destination: item.Destination,
		PersonPid:   publicid.Obfuscate32bit(item.PersonId),
		Shortmedium: item.Shortmedium,
		Shortobject: item.Shortobject,
	}

	return res
}
