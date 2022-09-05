package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main(){
	// 定义一个成员(44)5行4列的float数组，把里面的元素全部放到一个一维切片里。
	var floats  [5][4]float32
	s := []float32{}
	for _, items := range floats{
		for _, item := range items{
			s = append(s, item)
		}
	}
	
	// arr:=make([]int,3,4);brr:=append(arr,1)问arr里的元素是什么？arr和brr会相互影响吗？写代码验证一下。
	arr := make([]int, 3, 4)
	brr := append(arr, 1)
	for i, item := range arr{
		fmt.Printf("第%d位：%d\n", i, item)
	}
	brr[0] = 1
	fmt.Println(arr[0])
	fmt.Println("arr与brr回互相影响")

	// 创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
	sliceInt := make([]int,0,10)
	for i := 0; i < 100; i++{
		sliceInt = append(sliceInt, rand.Intn(128))
	}
	maps := make(map[int]bool, 100)
	cnt := 0
	for _, item := range sliceInt{
		if maps[item] {
			cnt++
		}else{
			maps[item] = true
		}
	}
	fmt.Printf("互不相同的元素有%d个\n", 100 - cnt)

	fmt.Println(arr2string([]int{}))
}

// 实现一个函数func arr2string(arr []int) string，比如输入[]int{2,4,6}，返回“2 4 6”。输入的切片可能很短，也可能很长。
func arr2string(arr []int)string{
	if arr == nil || len(arr) == 0{
		return ""
	}
	s := strings.Builder{}
	i := 0
	for i < len(arr)-1{
		c := strconv.Itoa(arr[i])
		s.WriteString(c)
		s.WriteString(" ")
		i++
	}
	s.WriteString(strconv.Itoa(arr[i]))
	return s.String()
}