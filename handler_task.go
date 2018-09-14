package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasksHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	c.JSON(http.StatusOK, OkData(GetTasks(pageParams.Page, pageParams.Size)))
}
