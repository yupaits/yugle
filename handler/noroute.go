package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yugle/common/result"
)

func NoRoute(c *gin.Context) {
	path := c.Request.URL.Path
	paths := strings.Split(path, "/")
	if paths[1] == "api" {
		c.JSON(http.StatusNotFound, result.CodeFail(result.NotFound))
	} else {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"path": path,
		})
	}
}
