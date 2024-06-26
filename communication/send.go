package communication

import (
	"github.com/we-and/weand_backend_common/app"
	m "github.com/we-and/weand_backend_common/models"
)

func Send(target string,r *app.RouteContext, pTeam *m.Team, pBatch *m.SendBatch, pubMap map[string]interface{}) {
	if (target=="ALL"){
		SendAll(r,pTeam,pBatch,pubMap)
	}
}


func SendAll(r *app.RouteContext, pTeam *m.Team, pBatch *m.SendBatch, pubMap map[string]interface{}) {
	if(pTeam!=nil){
		team:=(*pTeam)
			for _, s := range team.Seats {
				p := s.Person
				if p != nil {
					SendToPerson2(*r,team.ID, pBatch, p, pubMap)
				}
			}
	}
}
