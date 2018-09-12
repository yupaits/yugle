package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasksHandler(c *gin.Context) {
	var pageParams PageParams
	if c.ShouldBindQuery(&pageParams) == nil {
		if pageParams.Page == 0 {
			pageParams.Page = DefaultPage
		}
		if pageParams.Size == 0 {
			pageParams.Size = DefaultSize
		}
		c.JSON(http.StatusOK, OkData(GetTasks(pageParams.Page, pageParams.Size)))
	} else {
		c.JSON(http.StatusOK, CodeFail(ParamsError))
	}
}
