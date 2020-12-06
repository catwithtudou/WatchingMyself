package _7_facade

/**
 *@Author tudou
 *@Date 2020/12/6
 **/

type API interface {
	Test()string
}

func NewAPI()API{
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (impl *apiImpl)Test()string{
	a:=impl.a.TestA()
	b:=impl.b.TestB()
	return a+b
}

type AModuleAPI interface {
	TestA()string
}

func NewAModuleAPI()AModuleAPI{
	return &aModuleAPIImpl{}
}

type aModuleAPIImpl struct{}

func (aImpl *aModuleAPIImpl)TestA()string{
	return "A"
}

type BModuleAPI interface {
	TestB()string
}

func NewBModuleAPI()BModuleAPI{
	return &bModuleAPIImpl{}
}

type bModuleAPIImpl struct{}

func (bImpl *bModuleAPIImpl)TestB()string{
	return "B"
}