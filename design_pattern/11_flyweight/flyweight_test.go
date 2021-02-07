package _1_flyweight

import "testing"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/

func TestFlyweight(t *testing.T){
	viewer:=NewFileViewer("1.txt")
	viewer.PrintData()
}