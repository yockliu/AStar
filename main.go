package main

import (
	"fmt"
	"yockl/motodispatch/api"
	"yockl/motodispatch/roadmap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	maxX := float64(1000)
	maxY := float64(1000)
	rootSquare := roadmap.NewSquareTreeNode(
		roadmap.Point{X: 0, Y: 0},
		roadmap.Point{X: 0, Y: maxY},
		roadmap.Point{X: maxX, Y: maxY},
		roadmap.Point{X: maxX, Y: 0},
	)

	rootSquare.TreeQuarterCut(2)
	rootSquare.TreePrintLeaf()

	fmt.Println()

	// point := roadmap.Point{X: 999, Y: 999}
	// s := rootSquare.TreeFindLeafContains(&point)
	// roadmap.DebugPrint()
	// fmt.Println()
	// fmt.Println(s.Square)

	engine := gin.New()

	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	engine.Use(cors.New(corsConf))

	apiGroup := engine.Group("api")
	api.Init(apiGroup)

	engine.Run(fmt.Sprintf(":9000"))
}
