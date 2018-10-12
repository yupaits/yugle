package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routePage(template string, c *gin.Context) {
	c.HTML(http.StatusOK, template, gin.H{})
}
