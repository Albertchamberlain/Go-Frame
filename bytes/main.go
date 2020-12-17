package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 := "Φφϕ kKK"
	s2 := "ϕΦφ KkK"

	// 看看 s1 里面是什么
	for _, c := range s1 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s2 里面是什么
	for _, c := range s2 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	// 看看 s1 和 s2 是否相似
	fmt.Println(bytes.EqualFold([]byte(s1), []byte(s2)))

	//*****************************************************************

	bs := [][]byte{
		[]byte("Hello World !"),
		[]byte("Hello 世界！"),
		[]byte("hello golang ."),
	}
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！. "), r)
	}
	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimFunc(b, f))
	}
	// "Hello World"
	// "Hello 世界"
	// "Hello Golang"
	for _, b := range bs {
		fmt.Printf("%q\n", bytes.TrimPrefix(b, []byte("Hello ")))
	}
	// "World !"
	// "世界！"
	// "hello Golang ."

	s := "截取中文"
	//试试这样能不能截取?
	fmt.Println(s[:2])

	ss := "截取中文"
	//试试这样能不能截取? 将中文利用 [] rune 转换成 unicode 码点， 再利用 string 转化回去
	res := []rune(ss)
	fmt.Println(string(res[:2]))
}
