package main

import (
	"container/ring"
	"fmt"
)

type sumInt struct {
	Value int
}

//假设i为int
func (s *sumInt) add(i interface{}) {
	s.Value += i.(int)
}

func main() {
	r := ring.New(10)
	for i := 0; i < 10; i++ {
		r.Value = 1
		r = r.Next()
	}
	sum := sumInt{}
	r.Do(sum.add)
	fmt.Println(sum)
}
