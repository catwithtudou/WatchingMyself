package abstract_factory

import "fmt"

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

/**
抽象工厂模式用于生成系列产品的工厂，所生成的对象是有关联的。
若抽象工厂生成的对象无关联则退化成为工厂函数模式。
 */

type Men interface {
	Men()
}

type Women interface {
	Women()
}

//抽象工厂接口生产接口
type PeopleFactory interface {
	CreateMen()Men
	CreateWomen()Women
}

//Natural具体实现类
type NaturalMenImpl struct{}
type NaturalWomenImpl struct{}
func (m *NaturalMenImpl)Men(){
	fmt.Println("Natural men")
}
func (w *NaturalWomenImpl)Women(){
	fmt.Println("Natural women")
}
//Natural具体工厂类
type NaturalFactory struct {}
func (n *NaturalFactory)CreateMen()Men {
	return &NaturalMenImpl{}
}
func (n *NaturalFactory)CreateWomen()Women{
	return &NaturalWomenImpl{}
}

//Artificial具体实现类
type ArtificialMenImpl struct{}
type ArtificialWomenImpl struct{}
func (m *ArtificialMenImpl)Men(){
	fmt.Println("Artificial men")
}
func (w *ArtificialWomenImpl)Women(){
	fmt.Println("Artificial women")
}
//NArtificial具体工厂类
type ArtificialFactory struct {}
func (n *ArtificialFactory)CreateMen()Men{
	return &ArtificialMenImpl{}
}
func (n *ArtificialFactory)CreateWomen()Women{
	return &ArtificialWomenImpl{}
}
