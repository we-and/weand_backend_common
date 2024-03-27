package communication

import (
	//communication "github.com/we-and/weand_backend_common/communication"
	"errors"
	"fmt"
	"time"

	app "github.com/we-and/weand_backend_common/app"
	m "github.com/we-and/weand_backend_common/models"
	query "github.com/we-and/weand_backend_common/query"

	//	"gorm.io/gorm"

	//	project "github.com/we-and/weand_backend_common/project"
	structs "github.com/we-and/weand_backend_common/structs"
)

func SendNotification(r app.RouteContext, pBatch *m.SendBatch, pPerson *m.Person, pMessage *structs.NotificationMessage) (bool, error) {
	db := r.AppContext().GetDb()
	///read pointers
	if pMessage == nil {
		return false, errors.New("no message")
	}
	message := *pMessage
	if pPerson == nil {
		return false, errors.New("no person")
	}
	person := *pPerson

	title := message.Title
	content := message.Content
	personId := person.ID
	data := map[string]string{}
	success, err := SendNotif(r, personId, title, content, db, data)
	///save job status
	job := m.SendJob{
		PersonId:    person.ID,
		TeamId:      message.TeamId,
		RelatedId:   message.RelatedId,
		Destination: "app",
		Shortmedium: "N",
		Shortobject: message.GetShortObject(),
		Success:     success,
		Errdesc:     fmt.Sprintf("%v", err),
	}
	job.Started = true
	job.Sent = true
	t := time.Now()
	job.Senddate = &t

	if pBatch != nil {
		batch := (*pBatch)
		job.BatchId = (batch).ID
		(*pBatch).Jobs = append((*pBatch).Jobs, job)
	}
	if !query.Save(r, db, &job, "001") {
		return false, errors.New("job not saved")
	}

	return success, err

}
