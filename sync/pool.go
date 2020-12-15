package main

import (
	"fmt"
	"sync"
)

type item struct {
	value int
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return item{2}
		},
	}
	//pool.Put(item{})
	data := pool.Get()
	fmt.Println(data)
	pool.Put(item{})
	data2 := pool.Get()
	fmt.Println(data2)
	pool.Put(item{1})
	data3 := pool.Get()
	fmt.Println(data3)
	fmt.Println()
}

//看起来使用方式很简单，创建一个对象池的方式传进去一个New对象的函数，然后就是两个函数获取对象Get和放入对象Put。
