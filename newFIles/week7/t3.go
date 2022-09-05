package main

//改变切片中的元素（元素是结构体）。
type Student struct {
	weight float32 //体重，kg
	height float32 //身高，m
	bmi    float32 //weight/height^2
}
func CalBMI1(students []*Student) {
	for _, student := range students {
		student.bmi = student.weight / (student.height * student.height)
	}
}
func main3(){
	s := []*Student{{weight: 70, height: 1.83}, {weight:80, height: 1.79}}
	println(s[0].bmi)
	println(s[1].bmi)
	CalBMI1(s)
	println(s[0].bmi)
	println(s[1].bmi)
}