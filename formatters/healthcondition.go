package formatters

import (
	m "github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/publicid"
)

type HealthconditionFormattedItem struct {
	Pid  uint32 `json:"pid"`
	Name string `json:"name"`
	Idx  uint32 `json:"idx"`
}

func FormatHealthconditions(items []m.Healthcondition, langCode string) []HealthconditionFormattedItem {
	res := []HealthconditionFormattedItem{}
	for _, v := range items {
		res = append(res, FormatHealthcondition(v, langCode))
	}
	return res
}

func FormatHealthcondition(item m.Healthcondition, langCode string) HealthconditionFormattedItem {
	res := HealthconditionFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Idx:  item.Idx,
		Name: item.GetName(langCode),
	}
	return res
}
