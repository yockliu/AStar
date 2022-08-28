package roadmap

import "math"

type PointEntity interface {
	GetX() float64
	GetY() float64
	SetX(v float64)
	SetY(v float64)
}

type Point struct {
	X float64 `json:"x,omitempty" form:"x"`
	Y float64 `json:"y,omitempty" form:"y"`
}

func (p *Point) GetX() float64 {
	return p.X
}

func (p *Point) GetY() float64 {
	return p.Y
}

func (p *Point) SetX(v float64) {
	p.X = v
}

func (p *Point) SetY(v float64) {
	p.Y = v
}

func Mid(p1 Point, p2 Point) Point {
	return Point{
		X: (p1.X + p2.X) / 2,
		Y: (p1.Y + p2.Y) / 2,
	}
}

type Line struct {
	start Point
	end   Point
}

func NewLine(p1 Point, p2 Point) Line {
	var line Line
	if p2.X >= p1.X {
		line = Line{
			start: p1,
			end:   p2,
		}
	} else {
		line = Line{
			start: p2,
			end:   p1,
		}
	}
	return line
}

func valueIn(v float64, v1 float64, v2 float64) bool {
	if v1 < v2 {
		return v1 <= v && v <= v2
	} else if v2 < v1 {
		return v2 <= v && v <= v1
	} else {
		return math.Floor(v) == math.Floor(v1)
	}
}

func PointInLine(p *Point, p1 *Point, p2 *Point) bool {
	return valueIn(p.X, p1.X, p2.X) && valueIn(p.Y, p1.Y, p2.Y)
}

func (line *Line) kb() (bool, float64, float64) {
	var validK bool // k 是否有效，垂直则无效
	var k float64   // 斜率
	var b float64   // 在 y 轴上的位置

	deltaX := line.end.X - line.start.X
	deltaY := line.end.Y - line.start.Y

	if deltaX == 0 {
		validK = false
		k = math.MaxFloat64
		b = line.end.X
	} else {
		validK = true
		k = deltaY / deltaX
		b = line.end.Y - k*line.end.X
	}

	return validK, k, b
}

func CrossPoint(line1 Line, line2 Line) *Point {
	validK1, k1, b1 := line1.kb()
	validK2, k2, b2 := line2.kb()

	// 两条线都垂直
	if !validK1 && !validK2 {
		return nil
	}

	// line1 垂直
	if !validK1 {
		x := line1.start.X
		return &Point{
			X: x,
			Y: k2*x + b2,
		}
	}

	// line2 垂直
	if !validK2 {
		x := line2.start.X
		return &Point{
			X: x,
			Y: k1*x + b1,
		}
	}

	// 平行
	if k1 == k2 {
		return nil
	}

	x := (b2 - b1) / (k1 - k2)
	y := x*k1 + b1
	return &Point{
		X: x,
		Y: y,
	}

}

func Distance(p1 PointEntity, p2 PointEntity) float64 {
	return math.Sqrt(math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2))
}

func DistancePow(p1 PointEntity, p2 PointEntity) float64 {
	return math.Pow(p2.GetX()-p1.GetX(), 2) + math.Pow(p2.GetY()-p1.GetY(), 2)
}

func PerpendicularDistance(p PointEntity, p1 PointEntity, p2 PointEntity, foot PointEntity) float64 {
	A := p2.GetY() - p1.GetY()
	B := p1.GetX() - p2.GetX()
	C := p2.GetX()*p1.GetY() - p1.GetX()*p2.GetY()

	x := (B*B*p.GetX() - A*B*p.GetY() - A*C) / (A*A + B*B)
	y := (-A*B*p.GetX() + A*A*p.GetY() - B*C) / (A*A + B*B)
	d := math.Abs(A*p.GetX()+B*p.GetY()+C) / math.Sqrt(math.Pow(A, 2)+math.Pow(B, 2))

	foot.SetX(x)
	foot.SetY(y)
	return d
}
