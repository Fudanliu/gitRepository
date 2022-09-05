package main

import "math/rand"

//随机初始化两个8*5的矩阵，求两个矩阵的和（逐元素相加）

func main(){
	a := [8][5]int{}
	b := [8][5]int{}
	for i:=0; i < 8; i++{
		for j:= 0; j < 5; j++{
			a[i][j] = rand.Intn(65535)
			b[i][j] = rand.Intn(65535)
		}
	}
	
	c := [8][5]int{}
	for i:=0; i < 8; i++{
		for j:= 0; j < 5; j++{
			c[i][j] = a[i][j] + b[i][j]
		}
	}
}