package testing

import "yockl/motodispatch/roadmap"

var quarterRoot *roadmap.SquareTreeNode

func IndexMap(keypoints []*roadmap.KeyPoint) {
	quarterRoot = roadmap.NewSquareTreeNode(roadmap.Point{X: 0, Y: 0}, roadmap.Point{X: 0, Y: maxY}, roadmap.Point{X: maxX, Y: maxY}, roadmap.Point{X: maxX, Y: 0})
	quarterRoot.TreeQuarterCut(1)

	for _, kp := range keypoints {
		quarterRoot.TreeInsertKeyPoint(kp)
	}
}

func IndexMoto(motos []*roadmap.Moto) {
	for _, moto := range motos {
		quarterRoot.TreeInsertMoto(moto)
	}
}
