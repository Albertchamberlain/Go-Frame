package main

import (
	"bytes"
	"fmt"
)

func main() {
	b1 := []byte("Hello World!")
	b2 := []byte("Hello 世界！")
	buf := make([]byte, 6)
	rd := bytes.NewReader(b1)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "World!"

	rd.Reset(b2)
	rd.Read(buf)
	fmt.Printf("%q\n", buf) // "Hello "
	fmt.Printf("Size:%d, Len:%d\n", rd.Size(), rd.Len())
	// Size:15, Len:9

}
