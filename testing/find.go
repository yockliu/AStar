package testing

import (
	"fmt"
	"math"
	"yockl/motodispatch/roadmap"
)

func FindNearestMoto(position *roadmap.Point) (*Path, *roadmap.Moto) {
	node := quarterRoot.TreeFindKeypointsNode(position)
	fmt.Println("node")
	fmt.Println(node)
	kp := nearestKeypoint(position, node.KeyPoints)
	foot := findFoot(position, kp)

	fmt.Println("foot = ", foot)

	finder := NewFinder(foot)
	path := NewPath(finder, nil, foot)
	finder.insertSortedPath(path)
	return finder.findNearestMotoPath()
}

type Finder struct {
	usedKeyPoints map[string]*roadmap.KeyPoint
	sortedPaths   []*Path
}

func NewFinder(kp *roadmap.KeyPoint) *Finder {
	f := new(Finder)
	f.usedKeyPoints = make(map[string]*roadmap.KeyPoint)
	return f
}

func (f *Finder) used(kp *roadmap.KeyPoint) bool {
	return f.usedKeyPoints[kp.ID] != nil
}

func (f *Finder) use(kp *roadmap.KeyPoint) {
	f.usedKeyPoints[kp.ID] = kp
}

func (f *Finder) removeSortedPath(path *Path) {
	for i, p := range f.sortedPaths {
		if p == path {
			f.sortedPaths = append(f.sortedPaths[:i], f.sortedPaths[i+1:]...)
		}
	}
}

func (f *Finder) insertSortedPath(path *Path) {
	for i, p := range f.sortedPaths {
		if path.Distance < p.Distance {
			end := append([]*Path{}, f.sortedPaths[i:]...)
			begin := append(f.sortedPaths[0:i], path)
			f.sortedPaths = append(begin, end...)
			return
		}
	}
	f.sortedPaths = append(f.sortedPaths, path)
}

func (finder *Finder) findNearestMotoPath() (*Path, *roadmap.Moto) {
	for len(finder.sortedPaths) > 0 {
		path := finder.sortedPaths[0]
		currentKP := path.KeyPoint

		if len(path.KeyPoint.Motos) > 0 {
			moto := nearestMoto(&currentKP.Point, currentKP.Motos)
			return path, moto
		} else {
			if len(currentKP.Friends) == 0 {
				return nil, nil
			}

			finder.removeSortedPath(path)
			for _, fKeyPoint := range currentKP.Friends {
				if finder.used(fKeyPoint) {
					continue
				}
				newPath := NewPath(path.finder, path, fKeyPoint)
				finder.insertSortedPath(newPath)
				finder.use(fKeyPoint)
			}
		}
	}
	return nil, nil
}

type Path struct {
	finder   *Finder           `json:"-"`
	parent   *Path             `json:"-"`
	KeyPoint *roadmap.KeyPoint `json:"points,omitempty"`
	Distance float64           `json:"distance,omitempty"`
}

func NewPath(finder *Finder, parent *Path, kp *roadmap.KeyPoint) *Path {
	path := new(Path)
	path.finder = finder
	path.parent = parent
	path.KeyPoint = kp
	if parent == nil {
		path.Distance = 0
	} else {
		path.Distance = parent.Distance + roadmap.Distance(kp, parent.KeyPoint)
	}
	return path
}

func (path *Path) GetKeyPoints() []*roadmap.KeyPoint {
	kps := append([]*roadmap.KeyPoint{}, path.KeyPoint)
	for path.parent != nil {
		path = path.parent
		kps = append(kps, path.KeyPoint)
	}
	return kps
}

func nearestKeypoint(p *roadmap.Point, keyPoints []*roadmap.KeyPoint) *roadmap.KeyPoint {
	var min *roadmap.KeyPoint
	for _, kp := range keyPoints {
		if min == nil {
			min = kp
		} else {
			d1 := roadmap.Distance(&min.Point, p)
			d2 := roadmap.Distance(&kp.Point, p)
			fmt.Println("nearestKeypoint d1", min.Point, " ", d1)
			fmt.Println("nearestKeypoint d2", kp.Point, " ", d2)
			if d2 < d1 {
				min = kp
			}
		}
	}
	fmt.Println("nearestKeypoint min", min)
	return min
}

func nearestMoto(p *roadmap.Point, motos []*roadmap.Moto) *roadmap.Moto {
	var min *roadmap.Moto
	for _, moto := range motos {
		if min == nil {
			min = moto
		} else {
			d1 := roadmap.Distance(&min.Point, p)
			d2 := roadmap.Distance(&moto.Point, p)
			if d2 < d1 {
				min = moto
			}
		}
	}
	return min
}

func findFoot(p *roadmap.Point, keypoint *roadmap.KeyPoint) *roadmap.KeyPoint {
	var minDistance = math.MaxFloat64
	var minFoot *roadmap.KeyPoint
	p1 := keypoint
	for _, p2 := range keypoint.Friends {
		foot := new(roadmap.KeyPoint)
		foot.Friends = append(foot.Friends, p1)
		foot.Friends = append(foot.Friends, p2)
		distance := roadmap.PerpendicularDistance(p, p1, p2, foot)
		fmt.Println("findFoot p1 = ", p1, ", p2 = ", p2)
		fmt.Println("findFoot ", foot, ", distance = ", distance)
		if !roadmap.PointInLine(&foot.Point, &p1.Point, &p2.Point) {
			fmt.Println("not in ???")
			continue
		}
		if distance < minDistance {
			minDistance = distance
			minFoot = foot
		}
	}
	return minFoot
}
