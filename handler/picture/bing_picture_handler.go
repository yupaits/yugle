package picture_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yugle/common/pagination"
	"yugle/common/result"
	"yugle/crawler/picture"
)

func GetBingPicturesHandler(c *gin.Context) {
	var pageParams pagination.PageParams
	if c.ShouldBindQuery(&pageParams) == nil {
		if pageParams.Page == 0 {
			pageParams.Page = pagination.DefaultPage
		}
		if pageParams.Size == 0 {
			pageParams.Size = pagination.DefaultSize
		}
		c.JSON(http.StatusOK, result.OkData(crawler.GetBingPictures(pageParams.Page, pageParams.Size)))
	} else {
		c.JSON(http.StatusOK, result.CodeFail(result.ParamsError))
	}
}
