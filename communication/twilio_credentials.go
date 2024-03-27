package communication

import (
	//	"github.com/we-and/weand_backend_common/communication"
	config "github.com/we-and/weand_backend_common/config"
)

func BuildTwilioCredentials(appConfig *config.AppConfig) TwillioCredentials {
	sid := appConfig.Twilio.Sid
	token := appConfig.Twilio.Token
	fromNumber := appConfig.Twilio.From
	emailApiKey := appConfig.Twilio.API_KEY

	return BuildTwilio(sid, token, fromNumber, emailApiKey)
	//return BuildTwilio(appConfig.Twilio.Sid, appConfig.Twilio.Token, appConfig.Twilio.From, appConfig.Twilio.API_KEY)
}
