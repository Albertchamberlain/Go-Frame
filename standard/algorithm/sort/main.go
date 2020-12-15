package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	age  int
}

func main() {
	data := []person{
		{"Alice", 20},
		{"Bov", 21},
		{"Amos", 21},
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].age < data[j].age
	})

	for _, each := range data {
		fmt.Println("Name:", each.Name, "Age:", each.age)
	}
}
