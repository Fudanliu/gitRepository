package main

import "fmt"
//对一个任意长度的byte切片，在尾部追加任意元素使切片的长度变成8的整位数。这个过程称为padding。要求拿到padding后的切片能够还原出原始的切片。
func Padding(originSlice []byte) ([]byte,int){
	l := len(originSlice)
	for i:= l; i % 8 != 0; i++{
		originSlice = append(originSlice, 0)
	}
	return originSlice, l
 }
func UnPadding(paddedSlice []byte, len int) []byte { 
	return paddedSlice[0:len]
}

func main(){
	ps, l := Padding([]byte{1,2,3,4})
	fmt.Printf("%v",UnPadding(ps,l))
}