package avl

func rotateRight(y *Node) *Node {
	x := y.Right
	T2 := x.Right

	y.Left = T2
	x.Right = y

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	x.Right = T2
	y.Left = x

	return y
}
