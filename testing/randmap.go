package testing

import (
	"math/rand"
	"time"
	"yockl/motodispatch/roadmap"
)

var maxX float64 = 1280
var maxY float64 = 600
var allKeyPoints []*roadmap.KeyPoint

func GetKeyPoints() []*roadmap.KeyPoint {
	return allKeyPoints
}

func Randmap() {
	rand.Seed(time.Now().Unix())

	basePoint := roadmap.NewKeyPoint(randInRange(20, 30), randInRange(20, 30))

	allKeyPoints = genKeyPoints(basePoint, maxX, maxY)
	// for _, p := range allKeyPoints {
	// 	fmt.Println(p)
	// }
}

func genKeyPoints(p *roadmap.KeyPoint, maxX float64, maxY float64) []*roadmap.KeyPoint {
	var allKeyPoints []*roadmap.KeyPoint
	current := p
	vLine := genUpLine(current, maxY)

	for i := 0; i < 20; i++ {
		allKeyPoints = append(allKeyPoints, vLine...)

		var right *roadmap.KeyPoint
		var lastRight *roadmap.KeyPoint
		var thisVLine []*roadmap.KeyPoint

		for j, lp := range vLine {
			if j == 0 {
				right = randRightPoint(lp, 100, 200)
				if right.X >= maxX {
					break
				}
				current = right
			} else {
				right = roadmap.NewKeyPoint(current.X, lp.Y)
			}

			roadmap.AddFriends(lp, right)
			if lastRight != nil {
				roadmap.AddFriends(lastRight, right)
			}
			lastRight = right

			thisVLine = append(thisVLine, right)
		}

		vLine = thisVLine
	}

	IndexMap(allKeyPoints)

	return allKeyPoints
}

func genUpLine(p *roadmap.KeyPoint, maxY float64) []*roadmap.KeyPoint {
	var upline []*roadmap.KeyPoint

	for i := 0; i < 20; i++ {
		upline = append(upline, p)
		next := randUpPoint(p, 100, 200)
		if next.Y > maxY {
			break
		}
		roadmap.AddFriends(p, next)
		p = next
	}

	return upline
}

func randRightPoint(p *roadmap.KeyPoint, minDistance float64, maxDistance float64) *roadmap.KeyPoint {
	distance := randInRange(minDistance, maxDistance)
	rightPoint := roadmap.NewKeyPoint(p.X+distance, p.Y)
	return rightPoint
}

func randUpPoint(p *roadmap.KeyPoint, minDistance float64, maxDistance float64) *roadmap.KeyPoint {
	distance := randInRange(minDistance, maxDistance)
	upPoint := roadmap.NewKeyPoint(p.X, p.Y+distance)
	return upPoint
}

func randInRange(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
