package v1

import (
	"github.com/gin-gonic/gin"
	"wusthelper-mp-gateway/library/ecode"
)

type undergradLoginReq struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func undergradLogin(c *gin.Context) {
	platform := getPlatform(c)
	req := new(undergradLoginReq)
	err := c.BindJSON(req)
	if err != nil {
		response(c, ecode.ParamWrong, nil)
		return
	}

	ctx := c.Request.Context()
	_, _, err = serv.UndergradLogin(&ctx, req.UserAccount, req.UserPassword, "oid", platform)
	if err != nil {
		return
	}

	response(c, ecode.OK, "token")
	c.Next()
}
