package _0_composite

import "testing"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/

func TestComposite(t *testing.T){
	root := NewComponent(CompositeNode,"root")
	kid1 := NewComponent(CompositeNode,"kid1")
	kid2 := NewComponent(CompositeNode,"kid2")
	kid3 := NewComponent(CompositeNode,"kid3")

	root.AddChild(kid1)
	root.AddChild(kid2)
	kid2.AddChild(kid3)

	root.DetailPrint("***")
}


func TestLeaf(t *testing.T){
	root := NewComponent(CompositeNode,"root")
	kid1 := NewComponent(LeafNode,"kid1")
	kid2 := NewComponent(LeafNode,"kid2")
	kid3 := NewComponent(LeafNode,"kid3")

	root.AddChild(kid1)
	root.AddChild(kid2)
	root.AddChild(kid3)

	root.DetailPrint("***")
}