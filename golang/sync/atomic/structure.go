package main


// 不同架构下相同的代码进行编译测试
// 观察其编译指令是否为原子操作

const x int64 = 1 + 1<<33

func main(){
	var i = x
	_ = i
}
