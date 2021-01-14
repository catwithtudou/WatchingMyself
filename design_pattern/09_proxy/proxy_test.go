package _9_proxy

import "testing"

/**
 *@Author tudou
 *@Date 2021/1/14
 **/


func TestProxy_Do(t *testing.T) {
	var target Target
	target=&Proxy{}

	target.Do()
}