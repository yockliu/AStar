package api

import (
	"fmt"
	"net/http"
	"yockl/motodispatch/roadmap"
	"yockl/motodispatch/testing"

	"github.com/gin-gonic/gin"
)

func randomMoto(c *gin.Context) {
	testing.RandomMoto(10)
	all := testing.GetMotos()
	c.JSON(http.StatusOK, all)
}

func findNearestMoto(c *gin.Context) {
	point := new(roadmap.Point)
	c.ShouldBindQuery(point)
	fmt.Println("findNearestMoto point ", point)
	path, moto := testing.FindNearestMoto(point)
	motopath := new(MotoPath)
	if path != nil && moto != nil {
		motopath.Path = path.GetKeyPoints()
		motopath.Moto = moto
	}
	c.JSON(http.StatusOK, motopath)
}

type MotoPath struct {
	Moto *roadmap.Moto       `json:"moto,omitempty"`
	Path []*roadmap.KeyPoint `json:"path,omitempty"`
}
