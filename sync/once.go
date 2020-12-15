package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	tesatbody := func() { fmt.Println("just once") }
	for i := 0; i < 10; i++ {
		once.Do(tesatbody)
	}

}
