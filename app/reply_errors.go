package app

import (

	///"github.com/we-and/weand_backend_common/log"

	"github.com/we-and/weand_backend_common/response"
	//"github.com/we-and/weand_backend_common/log"
)

func SetBadRequest(r RouteContext, trigger string, err error, errCode string) {
	response.SetBadRequest(r.FiberCtx, trigger, err, errCode)
}

func OK(r RouteContext, obj interface{}) error {
	response.SetOK(r.FiberCtx, obj)
	return nil
}
func OKXML(r RouteContext, obj interface{}) error {
	response.SetOKXML(r.FiberCtx, obj)
	return nil
}

func OKHTML(r RouteContext, obj string) error {
	response.SetOKHTML(r.FiberCtx, obj)
	return nil
}

func SetNotFoundError(r RouteContext, trigger string, err error, errCode string) error {
	response.SetInternalError(r.FiberCtx, trigger, err, errCode)
	return nil
}
func SetInternalError(r RouteContext, trigger string, err error, errCode string) error {
	response.SetInternalError(r.FiberCtx, trigger, err, errCode)
	return nil
}

func SetError(r RouteContext, trigger string, err error, errortype string, errCode string) {
	switch errortype {
	case "INTERNALERROR":
		SetInternalError(r, trigger, err, errCode)
	case "UNAUTHORIZED":
		SetUnauthorized(r, trigger, err, errCode)
	}
	SetInternalError(r, trigger, err, errCode)
}
func SetUnauthorized(r RouteContext, trigger string, err error, errCode string) {
	response.SetInternalError(r.FiberCtx, trigger, err, errCode)
}
