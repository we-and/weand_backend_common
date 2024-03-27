package formatters

import (
	"time"

	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
	structs "github.com/we-and/weand_backend_common/structs"
	timezone "github.com/we-and/weand_backend_common/timezone"
)

type ChallengeFormattedItem struct {
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Pid  uint32             `json:"pid"`
	Name string `json:"name"`
	IsActive bool  `json:"is_active"`
}

func FormatChallenges(items []m.Challenge, me structs.Me, reqTzData *timezone.TzData) []ChallengeFormattedItem {
	res := []ChallengeFormattedItem{}
	for _, v := range items {
		res = append(res, FormatChallenge(v))
	}
	return res
}
func FormatChallenge(item m.Challenge) ChallengeFormattedItem {

	res := ChallengeFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
Name: item.Name,
IsActive: item.IsActive,
StartDate: item.StartDate,
EndDate: item.EndDate,

	}

	return res
}
