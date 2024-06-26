package auth

import (
	"github.com/we-and/weand_backend_common/app"
	"github.com/we-and/weand_backend_common/models"
	"github.com/we-and/weand_backend_common/query"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUserPublicIdAndCreate(r app.RouteContext, tx *gorm.DB, userId uint32) (resSuccess bool, resPublicid models.PublicId) {
	piduuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	pid := piduuid.String()

	//CREATE
	resPublicid = models.PublicId{UserID: userId, Publickey: pid}
	if !query.CreateOrRollback(r, tx, &resPublicid, "ME040-302") {
		return
	}
	resSuccess = true
	return
}

func GenerateUserPublicId(userId uint32) (resSuccess bool, resPublicid models.PublicId) {
	piduuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	pid := piduuid.String()

	//CREATE
	resPublicid = models.PublicId{UserID: userId, Publickey: pid}
	resSuccess = false
	return
}
