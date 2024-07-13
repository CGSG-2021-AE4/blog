package midleware

import (
	"strings"

	"github.com/CGSG-2021-AE4/blog/api"

	"github.com/gin-gonic/gin"
)

func AuthHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authorized", "false")
		c.Set("authErr", "")

		authStr := c.GetHeader("Authorization")
		if authStr == "" {
			c.Set("authErr", "no 'Authorization' header")
			return
		}
		strs := strings.Split(authStr, " ")
		if len(strs) != 2 {
			c.Set("authErr", "invalid auth str")
			return
		}
		if strs[0] != "Bearer" {
			c.Set("authErr", "invalid auth method")
			return
		}

		claims, err := us.ValidateToken(c, api.Token(strs[1]))
		if err != nil {
			c.Set("authErr", err.Error())
			return
		}

		c.Set("authorized", "true")
		c.Set("authId", claims.Issuer)
	}
}
