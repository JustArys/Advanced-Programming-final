package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller Controller) Index(c *gin.Context) {

	pd := controller.DefaultPageData(c)
	pd.Title = pd.Trans("Home")
	c.HTML(http.StatusOK, "index.html", pd)

}
