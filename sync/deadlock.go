package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	mutex.Lock()
	fmt.Println("Locked")
	mutex.Lock()
}
