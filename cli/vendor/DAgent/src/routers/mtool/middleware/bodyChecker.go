package middleware

import (
	"A-module/log"
	"A-module/routers/mtool/model"
	"DAgent/src/routers/mtool/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
)

func CheckBody(ctx *gin.Context) {
	body := model.Request{}

	if "GET" != ctx.Request.Method {
		err := ctx.ShouldBindBodyWith(&body, binding.JSON)
		if err != nil && err != io.EOF {
			log.Infof("Request Body Error : %v", err)
			api.BadRequest(ctx, model.Response{}, 10310)
			return
		}
	}
}
