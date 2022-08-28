package api

import (
	"github.com/gin-gonic/gin"
)

func Init(group *gin.RouterGroup) {
	group.GET("randomMap", randomMap)
	group.GET("randomMotos", randomMoto)
	group.GET("findNearestMoto", findNearestMoto)
}
