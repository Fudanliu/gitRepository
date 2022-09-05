package main

import (
	"fmt"
	"time"
)

//1. 把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func main1(){
	s1 := "1998-10-01 08:10:00"
	TIME_FMT := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TIME_FMT, s1, loc)
	fmt.Println(t)
	NEW_FMT := "20060102150405"
	s2 := t.Format(NEW_FMT)
	fmt.Println(s2)
}