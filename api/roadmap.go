package api

import (
	"net/http"
	"yockl/motodispatch/testing"

	"github.com/gin-gonic/gin"
)

func randomMap(c *gin.Context) {
	testing.Randmap()
	all := testing.GetKeyPoints()
	c.JSON(http.StatusOK, all)
}
