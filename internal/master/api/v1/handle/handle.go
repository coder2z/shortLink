package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortLink/internal/master/error/httpError"
	_map "shortLink/internal/master/map"
	"shortLink/internal/master/server"
	"shortLink/pkg/log"
	R "shortLink/pkg/response"
	"shortLink/pkg/validator"
)

type Handle struct{}

func (Handle) AddLink(ctx *gin.Context) {
	var link _map.AddLink
	if err := ctx.ShouldBind(&link); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(link); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if err := server.AddUrl(link); err != nil {
		log.Error(err.Error())
		R.Error(ctx, "error", nil)
	} else {
		R.Ok(ctx, "", nil)
	}
	return
}

func (Handle) Get(ctx *gin.Context) {
	var id _map.IdMap
	if err := ctx.ShouldBindUri(&id); err != nil {
		httpError.HandleBadRequest(ctx, nil)
		return
	}
	if err := validator.Struct(id); err != nil {
		httpError.HandleBadRequest(ctx, validator.GetMsg(err))
		return
	}
	if url, err := server.GetUrlByKey(id); err != nil {
		log.Error(err.Error())
		R.Error(ctx, "error", nil)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, url)
	}
	return
}
