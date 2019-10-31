package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/ToPeas/go-curd-template/mysql"
	"github/ToPeas/go-curd-template/pkg/app"
	"github/ToPeas/go-curd-template/pkg/e"
	"github/ToPeas/go-curd-template/pkg/util"
	"net/http"
)

type loginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	appG := app.Gin{Context: ctx}

	var input loginForm

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		appG.Error(http.StatusBadRequest, e.InvalidParams, err.Error())
		return
	}

	admin, err := mysql.GetAdminByUsername(input.Username)
	if err != nil {
		appG.Error(http.StatusOK, e.ERROR, err.Error())
		return
	}

	if admin == nil {
		appG.Error(http.StatusOK, e.UserNotFound, fmt.Errorf("%w username: %s", e.ErrUserNotFound, input.Username).Error())
		return
	}

	// 开始对比密码
	// 如果未加密过
	if admin.Password == input.Password {
		// 开始加密并保存
		hashed, err := util.HashPassword(admin.Password)
		if err != nil {
			appG.Error(http.StatusOK, e.ERROR, err.Error())
			return
		}

		// 存到db
		err = mysql.UpdatePasswordByUid(admin.ID, hashed)
		if err != nil {
			appG.Error(http.StatusOK, e.ERROR, err.Error())
			return
		}

		// 返回jwt
		jwt, err := util.GenerateToken(admin.ID)
		if err != nil {
			appG.Error(http.StatusOK, e.ERROR, err.Error())
			return
		}

		appG.Success(map[string]interface{}{
			"token": jwt,
		})
		return
	} else {
		// 对比 hash
		if util.CheckPasswordHash(input.Password, admin.Password) {
			// 正确
			// 返回jwt
			jwt, err := util.GenerateToken(admin.ID)
			if err != nil {
				appG.Error(http.StatusOK, e.ERROR, err.Error())
				return
			}

			appG.Success(map[string]interface{}{
				"token": jwt,
			})
			return
		} else {
			// 错误
			appG.Error(http.StatusOK, e.WrongPassword, "")
			return
		}
	}
}
