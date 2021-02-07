package _6_singleton

import "sync"

/**
 *@Author tudou
 *@Date 2020/12/6
 **/

type Singleton struct{}

var once sync.Once
var singleton *Singleton

//使用懒惰模式的单例模式，使用双重检查加锁保证线程安全
func GetInstance()*Singleton{
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}