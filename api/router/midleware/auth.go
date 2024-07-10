package midleware

import (
	"strings"

	"github.com/CGSG-2021-AE4/blog/api"

	"github.com/gin-gonic/gin"
)

func AuthHandler(us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authorized", false)

		authStr := c.GetHeader("Authorization")
		//log.Println("MIDLEWARE: ", authStr)
		strs := strings.Split(authStr, " ")
		if len(strs) != 2 {
			// log.Println("AUTH ERROR: invalid auth str")
			c.Set("authErr", "invalid auth str")
			return
		}
		if strs[0] != "Bearer" {
			//log.Println("AUTH ERROR: invalid auth method")
			c.Set("authErr", "invalid auth method")
			return
		}

		claims, err := us.ValidateToken(c, api.Token(strs[1]))
		if err != nil {
			// log.Println("AUTH ERROR:", err.Error())
			c.Set("authErr", err.Error())
			return
		}
		c.Set("authorized", true)
		c.Set("username", claims.Issuer)
	}
}
