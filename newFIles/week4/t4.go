package main

import "errors"
//实现函数func square(num interface{}) interface{}，计算一个interface{}的平方，interface{}允许是4种类型：float32、float64、int、byte
func main(){

}

func square(num interface{}) interface{}{
	switch num.(type){
	case float32:
		return num.(float32)*num.(float32)
	case float64:
		return num.(float64)*num.(float64)
	case int:
		return num.(int)*num.(int)
	case byte:
		return num.(byte)*num.(byte)
	default:
		return errors.New("非允许类别")
	} 
}