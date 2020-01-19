package middleware

import (
	"github.com/PaserDark/go_web/handler"
	"github.com/PaserDark/go_web/pkg/errno"
	"github.com/PaserDark/go_web/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}