package tree

/**
 *@Author tudou
 *@Date 2020/12/10
 **/

type Tree struct{
	Value int
	Left,Right *Tree
}

func NewNode(value int)*Tree {
	return &Tree{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func Pre(){


}