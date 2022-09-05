//hello包
package main

import (
	"reflect"
	"fmt"
)

type A struct {
	sex    bool
	age    int64
	weight uint16
}
type B struct {
	weight uint16
	sex    bool
	age    int64
}

func main() {
	v := make([]reflect.Value, 2)
	v[0] = reflect.ValueOf(1)
	v[1] = reflect.ValueOf("test")
	for _, ele := range v{
		fmt.Println(ele)
	}

}
