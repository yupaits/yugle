package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBingPicturesHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	c.JSON(http.StatusOK, OkData(GetBingPictures(pageParams.Page, pageParams.Size)))
}
