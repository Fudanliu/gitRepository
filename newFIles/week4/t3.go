package main

//定义两个接口：鱼类和爬行动物，再定义一个结构体：青蛙，同时实现上述两个接口

type fish interface{
	swim()
}
type reptile interface{
	move()
}

type frog struct{
	name string
}

func ( f frog) swim(){
	print("i can swim!")
}
func (f frog) move(){
	print("i can move~")
}

func main(){
	f := frog{"小青蛙"}
	f.swim()
	f.move()
}