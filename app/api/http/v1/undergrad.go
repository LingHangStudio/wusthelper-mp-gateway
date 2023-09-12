package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wusthelper-mp-gateway/library/ecode"
)

type undergradLoginReq struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func undergradLogin(c *gin.Context) {
	//platform := getPlatform(c)
	req := new(undergradLoginReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, apiResp[interface{}]{
			code: ecode.ParamWrong.Code(),
			msg:  "参数错误",
		})
	}

	c.JSON(http.StatusOK, apiResp[interface{}]{})
}
