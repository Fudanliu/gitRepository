package main

import "fmt"

////创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
type student struct{
	name string
	Cn,En,Math int
}


func main(){
	s := []student{{"刘备",80,80,80},{"关羽",100,70,70},{"张飞",90,60,75}}
	average(s[1])
	allAverage(s)
	count(s)
}

func average(s student){
	fmt.Printf("学生%s平均分为：%.2f\n", s.name, (float64(s.Cn+s.En+s.Math))/3)
}

func allAverage(s []student){
	cnt := 0
	var div float64 = float64(len(s) * 3)
	for _, ele := range s{
		cnt += ele.Cn
		cnt += ele.En
		cnt += ele.Math
	}
	total := float64(cnt)
	fmt.Printf("全班平均分为%.2f\n", total/div)
}

func count(s []student){
	cnt := 0
	for _, s := range s{
		if float64(s.Cn+s.En+s.Math)/3 < 60{
			cnt++
		}
	}
	fmt.Printf("全班一共%d个平均分不及格的\n",cnt)
}