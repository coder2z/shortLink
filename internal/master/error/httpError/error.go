package httpError

import (
	"shortLink/pkg/log"
	R "shortLink/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleInternalError(ctx *gin.Context) {
	log.Error("Internal Error", ctx.Err())
	R.Response(ctx, R.InternalError, "Internal Error", nil, http.StatusInternalServerError)
	return
}

func HandleBadRequest(ctx *gin.Context, data interface{}) {
	R.Response(ctx, R.BadRequest, "Bad Request", data, http.StatusBadRequest)
	return
}

func HandleNotFound(ctx *gin.Context) {
	R.Response(ctx, R.NotFound, "Not Found", nil, http.StatusNotFound)
	return
}

func HandleForbidden(ctx *gin.Context) {
	R.Response(ctx, R.Forbidden, "Forbidden", nil, http.StatusForbidden)
	return
}
