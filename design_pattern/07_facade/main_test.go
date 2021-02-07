package _7_facade

import (
	"fmt"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/12/6
 **/


func TestNewAPI(t *testing.T) {
	api:=NewAPI()
	result:=api.Test()
	fmt.Print(result)
}