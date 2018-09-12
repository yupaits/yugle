package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NoRoute(c *gin.Context) {
	path := c.Request.URL.Path
	paths := strings.Split(path, "/")
	if paths[1] == "api" {
		c.JSON(http.StatusNotFound, CodeFail(NotFound))
	} else {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"path": path,
		})
	}
}
