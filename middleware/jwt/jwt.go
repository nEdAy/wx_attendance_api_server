package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nEdAy/wx_attendance_api_server/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.Query("token")
		if token == "" {
			code = http.StatusUnauthorized // http.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized // http.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusUnauthorized // http.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  http.StatusText(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
