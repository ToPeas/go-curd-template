package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Context *gin.Context
}

func (g *Gin) Success(data interface{}) {
	g.Context.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"msg":     "",
		"data":    data,
	})

	return
}

func (g *Gin) Error(httpStatusCode int, msg string, hint string) {
	g.Context.JSON(httpStatusCode, map[string]interface{}{
		"success": false,
		"msg":     msg,
		"hint":    hint,
		"data":    nil,
	})

	return
}
