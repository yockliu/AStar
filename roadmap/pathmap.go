package roadmap

type Section struct {
	P1           *Point
	P2           *Point
	RunningDir   int
	RunningMotos []*Moto
	WatingMotos  []*Moto
}

type Path struct {
	Points []*Point
}

type PathMap struct {
	Paths []*Path
}

func NewPathMap() {
	path1 := new(Path)
	path1.Points = []*Point{{1, 1}, {2, 2}}
	path2 := new(Path)
	path2.Points = []*Point{{1, 1}, {2, 2}}
}
