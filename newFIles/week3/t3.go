package main

import "math/rand"

//有一个长度为20的int数组，分别用两个for循环求前半部分的和和后半部分的积，要求只遍历一次数组
func main(){
	a := [20]int{}
	for i:=0; i < 20; i++{
		a[i] = rand.Intn(10)
	}
	p := 1
	l := 1
	for i:=0; i < 10; i++{
		p *= a[i]
	}
	for i:=11; i < 20; i++{
		l *= a[i]
	}
}