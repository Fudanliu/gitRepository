package main
import	("time"
		"fmt")
//用defer优雅地打印函数的耗时
//在每一个return的地方打印函数的耗时，很麻烦
func foo(i int) int {
	begin := time.Now()
	defer func ()  {
		fmt.Printf("function use time %d ms\n", time.Since(begin).Milliseconds())
	}()
	if i < 10 {
		return i + 4
	} else if i < 20 {
		return i * 4
	} else {
		return 0
	}
}