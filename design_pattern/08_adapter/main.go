package _8_adapter

/**
 *@Author tudou
 *@Date 2020/12/6
 **/

type Target interface {
	Result()string
}

type Origin interface {
	Request()string
}

func NewOrigin()Origin{
	return &originImpl{}
}

type originImpl struct{}

func (a *originImpl)Request()string{
	return "test"
}

func NewAdapter(o Origin)Target{
	return &adapter{
		Origin:&originImpl{},
	}
}

type adapter struct{
	Origin
}

func (a *adapter)Result()string{
	return "no-test"
}


