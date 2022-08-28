package roadmap

import "fmt"

type Square struct {
	BL Point
	TL Point
	TR Point
	BR Point
}

type SquareTreeNode struct {
	Square
	Children  []*SquareTreeNode
	KeyPoints []*KeyPoint
	Motos     []*Moto
}

func (s *Square) IsContains(p *Point) bool {
	a := (s.TR.X-s.TL.X)*(p.Y-s.TL.Y) - (s.TR.Y-s.TL.Y)*(p.X-s.TL.X)
	b := (s.BR.X-s.TR.X)*(p.Y-s.TR.Y) - (s.BR.Y-s.TR.Y)*(p.X-s.TR.X)
	c := (s.BL.X-s.BR.X)*(p.Y-s.BR.Y) - (s.BL.Y-s.BR.Y)*(p.X-s.BR.X)
	d := (s.TL.X-s.BL.X)*(p.Y-s.BL.Y) - (s.TL.Y-s.BL.Y)*(p.X-s.BL.X)
	if (a >= 0 && b >= 0 && c >= 0 && d >= 0) || (a <= 0 && b <= 0 && c <= 0 && d <= 0) {
		return true
	}
	return false
}

func NewSquareTreeNode(p1 Point, p2 Point, p3 Point, p4 Point) *SquareTreeNode {
	node := new(SquareTreeNode)
	node.BL = p1
	node.TL = p2
	node.TR = p3
	node.BR = p4
	return node
}

func (node *SquareTreeNode) IsLeaf() bool {
	return len(node.Children) == 0
}

func (node *SquareTreeNode) AddChild(child *SquareTreeNode) {
	node.Children = append(node.Children, child)
}

func (node *SquareTreeNode) _quarterCut() {
	pointTop := Mid(node.TL, node.TR)
	pointLeft := Mid(node.TL, node.BL)
	pointBotton := Mid(node.BL, node.BR)
	pointRight := Mid(node.TR, node.BR)

	line1 := NewLine(pointLeft, pointRight)
	line2 := NewLine(pointTop, pointBotton)

	pointCenter := CrossPoint(line1, line2)

	if pointCenter == nil {
		panic("Square QuarterCut can't find Center Point")
	}

	node.AddChild(NewSquareTreeNode(node.BL, pointLeft, *pointCenter, pointBotton))
	node.AddChild(NewSquareTreeNode(pointLeft, node.TL, pointTop, *pointCenter))
	node.AddChild(NewSquareTreeNode(*pointCenter, pointTop, node.TR, pointRight))
	node.AddChild(NewSquareTreeNode(pointBotton, *pointCenter, pointRight, node.BR))
}

func (node *SquareTreeNode) TreeQuarterCut(level int) {
	if level > 0 {
		node._quarterCut()
		for _, child := range node.Children {
			child.TreeQuarterCut(level - 1)
		}
	}
}

func (node *SquareTreeNode) TreeInsertKeyPoint(kp *KeyPoint) {
	if node.IsContains(&kp.Point) {
		node.KeyPoints = append(node.KeyPoints, kp)
		for _, child := range node.Children {
			child.TreeInsertKeyPoint(kp)
		}
	}
}

func (node *SquareTreeNode) TreeInsertMoto(m *Moto) {
	if !node.IsContains(&m.Point) {
		node.Motos = append(node.Motos, m)
		for _, child := range node.Children {
			child.TreeInsertMoto(m)
		}
	}
}

func (node *SquareTreeNode) TreePrintLeaf() {
	if node.IsLeaf() {
		fmt.Println(node.Square)
	} else {
		for _, child := range node.Children {
			child.TreePrintLeaf()
		}
	}
}

func (node *SquareTreeNode) TreeFindKeypointsNode(p *Point) *SquareTreeNode {
	fmt.Println("TreeFindKeypointsNode: ", node)
	if !node.IsContains(p) || len(node.KeyPoints) == 0 {
		fmt.Println("not contain or no keypoints")
		return nil
	}

	if node.IsLeaf() {
		fmt.Println("is leaf")
		return node
	}

	for _, child := range node.Children {
		findedChild := child.TreeFindKeypointsNode(p)
		if findedChild != nil {
			fmt.Println("is leaf")
			return findedChild
		}
	}

	return node
}

func (node *SquareTreeNode) TreeFindLeaf(p *Point) *SquareTreeNode {
	if !node.IsContains(p) {
		return nil
	}

	if node.IsLeaf() {
		return node
	} else {
		for _, child := range node.Children {
			finded := child.TreeFindLeaf(p)
			if finded != nil {
				return finded
			}
		}
		return nil
	}
}
