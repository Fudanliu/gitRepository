package main


//给定月份，判断属于哪个季节。分别用if和switch实现
import (
	"bufio"
	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()

	//if 实现
	if s == "January" || s == "November" || s == "December"{
		print("冬天")
	}else if  s == "February" || s == "April" || s == "March"{
		print("春天")
	}else if  s == "June" || s == "May" || s == "July"{
		print("夏天")
	}else if  s == "August" || s == "September" || s == "October"{
		print("秋天")
	}else{
		print("输入的不是月份")
	}

	//switch 实现
	switch s{
	case "January", "November" ,  "December":print("冬天")
	case "February", "April", "March":print("春天")
	case "June" , "May" , "July":print("夏天")
	case "August" , "September", "October":print("秋天")
	default:print("输入的不是月份")
	}
}
