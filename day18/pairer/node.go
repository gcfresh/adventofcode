package pairer

import (
	"strconv"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func (n *Node) Copy() *Node {
	c := Node{
		Value: n.Value,
	}
	if n.Left != nil {
		c.Left = n.Left.Copy()
		c.Left.Parent = &c
	}
	if n.Right != nil {
		c.Right = n.Right.Copy()
		c.Right.Parent = &c
	}

	return &c
}

func (n *Node) GetMagnitude() int {
	m := 0

	if n.Left.IsValue() {
		m = m + 3*n.Left.Value
	} else {
		m = m + 3*n.Left.GetMagnitude()
	}

	if n.Right.IsValue() {
		m = m + 2*n.Right.Value
	} else {
		m = m + 2*n.Right.GetMagnitude()
	}

	return m
}

func (n *Node) GetValue() int {
	if n.Left == nil && n.Right == nil {
		return n.Value
	}
	return -1
}

func (n *Node) IsValue() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) GetString() string {
	s := ""
	s = s + n.getString()
	return s
}
func (n *Node) getString() string {
	s := ""
	if n.Left != nil {
		s = s + "[" + n.Left.getString()
	}
	if n.Right != nil {
		s = s + "," + n.Right.getString() + "]"
	}
	if n.IsValue() {
		s = s + strconv.Itoa(n.Value)
	}
	return s
}

func (n *Node) Reduce() {
	depth := 0
	more := true
	//var l, r int
	//fmt.Println("Reducing:", n.GetString())
	for more {
		mr := true
		for mr {
			mr, _, _ = n.reduceExplode(depth)
			//fmt.Println("l", l, "r", r, "more", mr, "reduced to:", n.GetString())
		}
		ms := true
		ms, _, _ = n.reduceSplit(depth)

		more = mr || ms
		//fmt.Println("split", "more", more, "reduced to:", n.GetString())
	}
}

//
//// done vs explode
//func (n *Node) reduce(depth int) (bool, int, int) {
//	depth++
//	changed := false
//	var addLeft, addRight int
//
//	// exploded
//	if depth > 4 {
//		if n.Left.IsValue() && n.Right.IsValue() {
//			l, r := n.explode()
//			//fmt.Println("explode", l, r)
//			return true, l, r
//		}
//	}
//	// split?
//	if !changed && n.Left.GetValue() > 9 {
//		fmt.Println("split", n.Left.GetValue())
//		n.Left.split()
//		//if depth == 4 {
//		//	addLeft, addRight = n.Left.explode()
//		//	if addRight > 0 {
//		//		if n.Right.addToLeftMostNode(addRight) {
//		//			return true, addLeft, 0
//		//		}
//		//	}
//		//}
//		return true, addLeft, addRight
//	}
//
//	if !changed && n.Right.GetValue() > 9 {
//		fmt.Println("split", n.Right.GetValue())
//		n.Right.split()
//		//if depth == 4 {
//		//	addLeft, addRight = n.Right.explode()
//		//	if addLeft > 0 {
//		//		if n.Left.addToRightMostNode(addLeft) {
//		//			return true, 0, addRight
//		//		}
//		//	}
//		//}
//		return true, addLeft, addRight
//	}
//	if !n.Left.IsValue() {
//		changed, addLeft, addRight = n.Left.reduce(depth)
//		if changed && addRight > 0 {
//			if n.Right.addToLeftMostNode(addRight) {
//				return changed, addLeft, 0
//			}
//		}
//	}
//
//	if !n.Right.IsValue() && !changed {
//		changed, addLeft, addRight = n.Right.reduce(depth)
//		if changed && addLeft > 0 {
//			if n.Left.addToRightMostNode(addLeft) {
//				return changed, 0, addRight
//			}
//		}
//	}
//
//	return changed, addLeft, addRight
//}

func (n *Node) reduceExplode(depth int) (bool, int, int) {
	//fmt.Println("HELLOOOoooooooooooooooooo")
	depth++
	changed := false
	var addLeft, addRight int

	// exploded
	if depth > 4 {
		if n.Left.IsValue() && n.Right.IsValue() {
			l, r := n.explode()
			//fmt.Println("explode [" + strconv.Itoa(l) + "," + strconv.Itoa(r) + "]")
			return true, l, r
		}
	}

	if !n.Left.IsValue() {
		changed, addLeft, addRight = n.Left.reduceExplode(depth)
		if changed && addRight > 0 {
			if n.Right.addToLeftMostNode(addRight) {
				return changed, addLeft, 0
			}
		}
	}

	if !n.Right.IsValue() && !changed {
		changed, addLeft, addRight = n.Right.reduceExplode(depth)
		if changed && addLeft > 0 {
			if n.Left.addToRightMostNode(addLeft) {
				return changed, 0, addRight
			}
		}
	}

	return changed, addLeft, addRight
}

func (n *Node) reduceSplit(depth int) (bool, int, int) {
	depth++
	changed := false
	var addLeft, addRight int

	// split?
	if n.Left.GetValue() > 9 {
		//fmt.Println("split", n.Left.GetValue())
		n.Left.split()
		//if depth == 4 {
		//	addLeft, addRight = n.Left.explode()
		//	if addRight > 0 {
		//		if n.Right.addToLeftMostNode(addRight) {
		//			return true, addLeft, 0
		//		}
		//	}
		//}
		return true, 0, 0
	} else if !n.Left.IsValue() {
		changed, addLeft, addRight = n.Left.reduceSplit(depth)
		if changed && addRight > 0 {
			if n.Right.addToLeftMostNode(addRight) {
				return changed, addLeft, 0
			}
		}
	}

	if !changed && n.Right.GetValue() > 9 {
		//fmt.Println("split", n.Right.GetValue())
		n.Right.split()
		//if depth == 4 {
		//	addLeft, addRight = n.Right.explode()
		//	if addLeft > 0 {
		//		if n.Left.addToRightMostNode(addLeft) {
		//			return true, 0, addRight
		//		}
		//	}
		//}
		return true, 0, 0
	} else if !n.Right.IsValue() && !changed {
		changed, addLeft, addRight = n.Right.reduceSplit(depth)
		if changed && addLeft > 0 {
			if n.Left.addToRightMostNode(addLeft) {
				return changed, 0, addRight
			}
		}
	}
	if !n.Left.IsValue() {

	}

	if !n.Right.IsValue() && !changed {

	}

	return changed, addLeft, addRight
}

func (n *Node) split() {
	n.Left = &Node{Value: n.Value / 2, Parent: n}
	n.Right = &Node{Value: n.Value - n.Value/2, Parent: n}
	n.Value = 0
}

func (n *Node) explode() (int, int) {
	l := n.Left.GetValue()
	r := n.Right.GetValue()
	n.Left = nil
	n.Right = nil
	n.Value = 0

	return l, r
}

func (n *Node) addToLeftMostNode(v int) bool {
	// adding left most right
	added := false
	if n.IsValue() {
		//fmt.Println("adding right", n.Value, v)
		n.Value = n.Value + v
		//if n.Value > 10 {
		//	n.explode()
		//}
		added = true
	} else {
		added = n.Left.addToLeftMostNode(v)
		if !added && !n.Right.IsValue() {
			added = n.Right.addToLeftMostNode(v)
		}
	}
	return added
}

func (n *Node) addToRightMostNode(v int) bool {
	// adding right most left
	added := false
	if n.IsValue() {
		//fmt.Println("adding left", n.Value, v)
		n.Value = n.Value + v
		//if n.Value > 10 {
		//	n.explode()
		//}
		added = true
	} else {
		added = n.Right.addToRightMostNode(v)
		if !added && !n.Left.IsValue() {
			added = n.Left.addToRightMostNode(v)
		}
	}
	return added
}
