package _0_composite

import "fmt"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(name string)
	AddChild(Component)
	DetailPrint(string)
}

const (
	LeafNode = iota + 1
	CompositeNode
)

func NewComponent(kind int,name string)Component{
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeft()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)

	return c
}

type component struct{
	parent Component
	name string
}

func (c *component)Parent()Component{
	return c.parent
}

func (c *component)SetParent(parent Component){
	c.parent=parent
}

func (c *component)Name()string{
	return c.name
}

func (c *component)SetName(name string){
	c.name=name
}

func (c *component)AddChild(Component){}

func (c *component)DetailPrint(string){}

//单节点
type Leaf struct{
	component
}

func NewLeft()*Leaf{
	return &Leaf{}
}

func (c *Leaf)DetailPrint(pre string){
	fmt.Printf("%s——%s\n",pre,c.Name())
}

//组合节点
type Composite struct{
	component
	kids []Component
}

func NewComposite()*Composite{
	return &Composite{
		kids:     make([]Component,0),
	}
}

func (c *Composite)AddChild(kid Component){
	kid.SetParent(c)
	c.kids=append(c.kids,kid)
}

func (c *Composite)DetailPrint(pre string){
	fmt.Printf("%s++%s\n",pre,c.Name())
	pre+=" "
	for _,v:=range c.kids{
		v.DetailPrint(pre)
	}
}
