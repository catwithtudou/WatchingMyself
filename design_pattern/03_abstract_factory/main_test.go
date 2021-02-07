package abstract_factory

import "testing"

/**
 *@Author tudou
 *@Date 2020/11/30
 **/

func getMenAndWomen(factory PeopleFactory) {
	factory.CreateMen().Men()
	factory.CreateWomen().Women()
}

func TestArtificialFactory(t *testing.T) {
	var factory PeopleFactory
	factory = &ArtificialFactory{}
	getMenAndWomen(factory)
}

func TestNaturalFactory(t *testing.T) {
	var factory PeopleFactory
	factory = &NaturalFactory{}
	getMenAndWomen(factory)
}