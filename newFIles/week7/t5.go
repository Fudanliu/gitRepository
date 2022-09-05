package main

type vehicle interface{
	scream()string
}
type bike struct{}
type car struct{}
type train struct{}

func (b bike)scream()string{
	return "叮铃铃"
}
func (c car)scream()string{
	return "嘟嘟嘟"
}
func (t train)scream()string{
	return "呜呜呜"
}
