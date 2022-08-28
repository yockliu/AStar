package testing

import (
	"math/rand"
	"yockl/motodispatch/roadmap"
)

var allMoto []*roadmap.Moto

func GetMotos() []*roadmap.Moto {
	return allMoto
}

func RandomMoto(count int) {
	allMoto = nil
	for i := 0; i < count; i++ {
		keypoint := randomKeyPoint(allKeyPoints)
		anotherKeypoint := randomKeyPoint(keypoint.Friends)
		p := randomPointInLine(&keypoint.Point, &anotherKeypoint.Point)
		moto := roadmap.NewMoto(p.X, p.Y)
		keypoint.Motos = append(keypoint.Motos, moto)
		anotherKeypoint.Motos = append(anotherKeypoint.Motos, moto)
		allMoto = append(allMoto, moto)
	}

	IndexMoto(allMoto)
}

func randomKeyPoint(kps []*roadmap.KeyPoint) *roadmap.KeyPoint {
	index := rand.Intn(len(kps))
	return kps[index]
}

func randomPointInLine(p1 *roadmap.Point, p2 *roadmap.Point) *roadmap.Point {
	deltaX := p2.X - p1.X
	deltaY := p2.Y - p1.Y

	if deltaX == 0 {
		x := p1.X
		y := p1.Y + deltaY*rand.Float64()
		return &roadmap.Point{
			X: x,
			Y: y,
		}
	} else {
		k := deltaY / deltaX
		distanceX := deltaX * rand.Float64()
		x := p1.X + distanceX
		y := p1.Y + k*distanceX
		return &roadmap.Point{
			X: x,
			Y: y,
		}
	}
}

func Sort(v1 float64, v2 float64) (float64, float64) {
	if v1 < v2 {
		return v1, v2
	} else {
		return v2, v1
	}
}
