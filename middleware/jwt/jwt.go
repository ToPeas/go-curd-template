package jwt

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github/ToPeas/go-curd-template/mysql"
	"github/ToPeas/go-curd-template/pkg/app"
	"github/ToPeas/go-curd-template/pkg/e"
	"github/ToPeas/go-curd-template/pkg/util"
	"log"

	"net/http"
	"strconv"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{Context: c}

		var msg string

		msg = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			msg = e.MissingJwt
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				log.Println(fmt.Errorf("%w token: %s", err, token))
				msg = e.AuthCheckTokenFail
			} else {
				uid := claims.Uid

				if uid <= 0 {
					msg = e.BrokenJwt
					appG.Error(http.StatusUnauthorized, msg, "")

					c.Abort()
					return
				}

				user, _ := mysql.GetAdminByUid(uid)
				if user == nil {
					msg = e.UserNotFound
					appG.Error(http.StatusUnauthorized, msg, fmt.Errorf("%w uid: %d", e.ErrUserNotFound, uid).Error())

					c.Abort()
					return
				}

				c.Set("uid", uid)
				raven.SetUserContext(&raven.User{
					ID: strconv.FormatInt(uid, 10),
					IP: c.ClientIP(),
				})
			}
		}

		if msg != e.SUCCESS {
			appG.Error(http.StatusUnauthorized, msg, "")

			c.Abort()
			return
		}

		c.Next()
	}
}
