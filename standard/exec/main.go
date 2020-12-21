package main

import (
	"fmt"
	"os/exec"
)

func main() {
	f, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}
