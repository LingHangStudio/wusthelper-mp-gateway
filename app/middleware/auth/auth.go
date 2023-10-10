package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wusthelper-mp-gateway/app/conf"
	api "wusthelper-mp-gateway/library/ecode/resp"
	_token "wusthelper-mp-gateway/library/token"
)

var (
	jwt *_token.Token
	dev bool
)

func Init(c *conf.Config) {
	jwt = _token.New(c.Server.TokenSecret, c.Server.TokenTimeout)
	dev = c.Server.Env == conf.DevEnv
}

func UserTokenCheck(c *gin.Context) {
	token := c.GetHeader("Token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": api.TokenInvalid,
			"msg":  "token invalid",
		})
		return
	}

	claims, valid := jwt.GetClaimVerify(token)
	if (!dev && !valid) || claims == nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": api.TokenInvalid,
			"msg":  "token invalid",
		})
		return
	}

	oid := (*claims)["oid"]
	c.Set("oid", oid)
	c.Next()
}
