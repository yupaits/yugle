package yugle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetShotOnOnePlusPicturesHandler(c *gin.Context) {
	pageParams := NewPageParams()
	c.ShouldBindQuery(pageParams)
	c.JSON(http.StatusOK, OkData(GetShotOnOnePlusPictures(pageParams.Page, pageParams.Size)))
}
