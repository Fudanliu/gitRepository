package main

import (
	"fmt"
	"time"
)

//2. 我们是每周六上课，输出我们未来4次课的上课日期（不考虑法定假日）
func main2(){
	dur := time.Duration(7* 24 * time.Hour)
	TIME_FMT := "2006-01-02 15:04:05"
	s := "2022-09-03 09:00:00"
	loc, _:=time.LoadLocation("Asia/Shanghai")
	t,_ := time.ParseInLocation(TIME_FMT, s, loc)
	for i:=0;i < 4; i++{
		t = t.Add(dur)
		fmt.Println(t)
	}
}