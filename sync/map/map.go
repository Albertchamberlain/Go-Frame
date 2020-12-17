package main

func main() {
	m := make(map[int]int)

	go func() {
		for {
			_ = m[1]
		}
	}()

	go func() {
		for {
			m[2] = 2
		}
	}()

	//select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行,这里是为了不让main退出
	select {}
}
