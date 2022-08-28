package roadmap

import "fmt"

var KeyPointIDCounter = 1

type KeyPoint struct {
	Point      `json:"point,omitempty"`
	ID         string      `json:"id,omitempty"`
	FriendsIDs []string    `json:"friendsIDs,omitempty"`
	Friends    []*KeyPoint `json:"-"`
	Motos      []*Moto     `json:"-"`
}

func NewKeyPoint(x float64, y float64) *KeyPoint {
	kp := new(KeyPoint)
	kp.X = x
	kp.Y = y
	kp.ID = fmt.Sprintf("%6d", KeyPointIDCounter)
	KeyPointIDCounter = KeyPointIDCounter + 1
	return kp
}

func AddFriends(p1 *KeyPoint, p2 *KeyPoint) {
	p1.Friends = append(p1.Friends, p2)
	p1.FriendsIDs = append(p1.FriendsIDs, p2.ID)
	p2.Friends = append(p2.Friends, p1)
	p2.FriendsIDs = append(p2.FriendsIDs, p1.ID)
}
