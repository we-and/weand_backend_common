package app

import (
	config "github.com/we-and/weand_backend_common/config"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"

	//"github.com/we-and/weand_backend_common/querier"
	"gorm.io/gorm"
)

type AppHandlerInterface interface {
	GetDb() *gorm.DB
	GetConfig() *config.AppConfig
	GetFirebase() *firebase.App
	GetGCP() *storage.Client
	//GetQuerier() app.Querier
}
