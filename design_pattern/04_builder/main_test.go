package builder

import (
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/11/30
 **/


func  TestProduct1(t *testing.T) {
	product1 := &Product1{}
	product := NewProduct(product1)
	product.WayMake()
	product1.Result()
}


func  TestProduct2(t *testing.T){
	product2 := &Product2{}
	product := NewProduct(product2)
	product.WayMake()
	product2.Result()
}