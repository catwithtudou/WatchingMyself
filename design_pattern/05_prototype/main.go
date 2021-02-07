package _5_prototype

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

/**
	原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。
	原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。
 */

type CloneType interface {
	Clone() CloneType
}

type PrototypeManager struct {
	prototypes map[string]CloneType
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]CloneType),
	}
}

func (p *PrototypeManager) Get(name string) CloneType {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype CloneType) {
	p.prototypes[name] = prototype
}