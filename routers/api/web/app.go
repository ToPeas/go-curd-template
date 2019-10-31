package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github/ToPeas/go-curd-template/mysql"
	"github/ToPeas/go-curd-templatemy/sql"
	"github/ToPeas/go-curd-templatepkg/app"
	"github/ToPeas/go-curd-templatepkg/e"
	"github/ToPeas/go-curd-templatepkg/myerr"
	"github/ToPeas/go-curd-templatepkg/validator"
	"net/http"
	"strconv"
)

type linkPayload struct {
	Name string  `form:"name" json:"name"  validate:"gt=6,lt=30,required" comment:"链接名称"`
	Url  string `form:"url" json:"url" validate:"gt=0,required" comment:"链接地址"`
}

// 获取所有的链接

func GetAllLinks(c *gin.Context) {
	appG := app.Gin{Context: c}

	links, err := mysql.DaoGetLinks()

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Success(links)
	return
}

// 获取单个link

func GetLink(c *gin.Context) {
	appG := app.Gin{Context: c}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appG.Error(http.StatusBadRequest, e.InvalidParams, err.Error())
		return
	}

	link, err := mysql.DaoGetLink(id)

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Success(link)
	return
}

// 添加链接

func AddLink(c *gin.Context) {
	appG := app.Gin{Context: c}

	var linkP linkPayload
	// 验证参数格式
	if err := c.ShouldBindJSON(&linkP); err != nil {
		// 当需要的参数类型不匹配的时候会出现错误，属于开发测试时就应该发现的异常
		str := myerr.NewParameterError("绑定参数异常")
		appG.Error(http.StatusBadRequest, str.Message, err.Error())
	}
	// 验证参数
	if err := validator.GlobalValidator.Check(&linkP); err != nil {
		str := myerr.NewParameterError(err.Error())
		appG.Error(http.StatusBadRequest, str.Message, err.Error())
		return
	}

	link := &mysql.Link{}

	_ = copier.Copy(link, &linkP)

	err := mysql.DaoAddLink(link)

	if err != nil {
		appG.Error(http.StatusOK, e.ERROR, err.Error())
		return
	}

	appG.Success(nil)
	return
}

// 更新Link
func PatchLink(c *gin.Context) {
	appG := app.Gin{Context: c}

	var linkP linkPayload

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		appG.Error(http.StatusBadRequest, e.InvalidParams, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&linkP); err != nil {
		// 当需要的参数类型不匹配的时候会出现错误，属于开发测试时就应该发现的异常
		str := myerr.NewParameterError("绑定参数异常")
		appG.Error(http.StatusBadRequest, str.Message, err.Error())
	}
	// 验证参数
	if err := validator.GlobalValidator.Check(&linkP); err != nil {
		str := myerr.NewParameterError(err.Error())
		appG.Error(http.StatusBadRequest, str.Message, err.Error())
		return
	}

	link := &mysql.Link{}
	_ = copier.Copy(link, &linkP)
	err = mysql.DaoUpdateLinkById(id, link)

	if err != nil {
		appG.Error(http.StatusOK, e.ERROR, err.Error())
		return
	}

	appG.Success(nil)
	return
}

// 删除Link

func DeleteLink(c *gin.Context) {
	appG := app.Gin{Context: c}
	var err error

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appG.Error(http.StatusBadRequest, e.InvalidParams, err.Error())
		return
	}

	err = mysql.DaoDeleteLink(id)
	if err != nil {
		appG.Error(http.StatusOK, e.ERROR, err.Error())
		return
	}

	appG.Success(nil)
	return
}

