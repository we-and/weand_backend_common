package formatters

import (
	"time"

	"github.com/we-and/weand_backend_common/publicid"

	m "github.com/we-and/weand_backend_common/models"
)

type PlanFormattedItem struct {
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	Pid       uint32                 `json:"pid"`
	Name      string                 `json:"name"`
	Pricings  []PricingFormattedItem `json:"pricings,omitempty"`
}

func FormatPlan(v m.Plan) PlanFormattedItem {

	res := PlanFormattedItem{
		CreatedAt: v.CreatedAt,
		Pid:       publicid.Obfuscate32bit(v.ID),
		Name:      v.Name,
		Pricings:  FormatPricings(v.Pricings),
	}
	return res
}

func FormatPlans(items []m.Plan) []PlanFormattedItem {
	res := []PlanFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPlan(v))
	}
	return res
}
