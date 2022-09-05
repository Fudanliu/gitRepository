package main

func beEqual(arr, brr [] byte)bool{
	if cap(arr)!=cap(brr) || len(arr) != len(brr){
		return false
	}
	for i:= 0; i < len(arr);i++{
		if arr[i] != brr[i]{
			return false
		}
	}
	return true
}