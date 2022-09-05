package main

import "errors"

//实现一个函数，接受若干个float64（用不定长参数），返回这些参数乘积的倒数，除数为0时返回error
func main(){

}

func f(floats...float64) (float64, error){
	total := 1.0
	for _, ele := range floats{
		total*=ele
	} 
	if total - float64(int(total)) < 0.0001{
		return 0, errors.New("divide 0")
	}
	return 1.0/total, nil
}