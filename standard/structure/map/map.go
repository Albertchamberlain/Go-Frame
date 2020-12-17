package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[string]int = map[string]int{"cmos": 12, "amos": 12, "albert": 10}
	i := m["amos"]
	fmt.Println(i)

	var keys []string

	// 把key单独抽取出来，放在数组中
	for k, _ := range m {
		keys = append(keys, k)
	}
	// 进行数组的排序
	sort.Strings(keys)
	// 遍历数组就是有序的了
	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}
