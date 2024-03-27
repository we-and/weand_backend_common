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

func SendSMSReceipt(r app.RouteContext, to string, pMessage *structs.SMSMessage) (bool, error) {
	if pMessage == nil {
		return false, errors.New("no message")
	}
	message := *pMessage

	///get config
	appConfig := r.AppCtx.GetConfig()
	credentials := BuildTwilioCredentials(appConfig)
	success, err := SendSMSWithClient(r, &credentials, to, message.Content)
	return success, err

}
func SendSMS(r app.RouteContext, pBatch *m.SendBatch, pPerson *m.Person, to string, pMessage *structs.SMSMessage) (bool, error) {
	///read pointers
	if pMessage == nil {
		return false, errors.New("no message")
	}
	message := *pMessage
	if pPerson == nil {
		return false, errors.New("no person")
	}
	person := *pPerson

	///get config
	appConfig := r.AppCtx.GetConfig()
	credentials := BuildTwilioCredentials(appConfig)

	//send
	content := fmt.Sprintf("%v %v", message.Title, message.Content)
	success, err := SendSMSWithClient(r, &credentials, to, content)

	///save job status
	job := m.SendJob{
		TeamId:      message.TeamId,
		PersonId:    person.ID,
		RelatedId:   message.RelatedId,
		Destination: to,
		Shortmedium: "S",
		Shortobject: message.GetShortObject(),
		Success:     success,
		Errdesc:     fmt.Sprintf("%v", err),
	}
	job.Started = true
	job.Sent = true
	t := time.Now()
	job.Senddate = &t
	if pBatch != nil {
		job.BatchId = (*pBatch).ID
		(*pBatch).Jobs = append((*pBatch).Jobs, job)

	}
	db := r.AppContext().GetDb()
	if !query.Save(r, db, &job, "001") {
		return false, errors.New("job not saved")
	}
	return success, err
}
