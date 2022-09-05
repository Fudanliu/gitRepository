package main


//定义结构体方法时，用结构体的值和指针有什么区别？用代码验证一下有没有对结构体进行拷贝


//用结构体的值是复制结构体内的所有变量，指针则是复制结构体地址
type user struct{
	name string
}
func (u user)changeName1(){
	u.name = "曹操"
}
func (u *user)changeName2(){
	u.name = "曹操"
}
func main(){
	u1 := user{"刘备"}
	u1.changeName1()
	println(u1.name)
	(&u1).changeName2()
	print(u1.name)

}