pacakge main

var t float64 = 1
//递归调用
func main(){
	
}

func f(floats ...float64) (float64, error) {
	if len(floats) == 0 {
		return 0, errors.New("Divisor of 0")
	}
	if floats[0] == 0 {
		return 0, errors.New("Divisor of 0")
	}
	t *= floats[0]
	if len(floats) == 1 {
		println(t)
		return 1 / t, nil
	}
	return f(floats[1:]...)
}

