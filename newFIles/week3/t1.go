package main

//实现一个结构体同时“继承”另外两个结构体
type dad struct{
	name,occupation string 
}
type mom struct{
	name,occupation string
}
type child struct{
	dad
	mom
	name string
}
func main(){
	c := child{dad{"刘备","皇帝"}, mom{"糜夫人","士族"}, "刘禅"}
	print(c)
}