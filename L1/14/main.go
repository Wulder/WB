package main

import (
	"fmt"
	"reflect"
)

func main() {

	var unknown interface{}

	unknown = "sss"
	fmt.Println(getType(unknown))
	unknown = 12
	fmt.Println(getType(unknown))
	unknown = true
	fmt.Println(getType(unknown))
	unknown = make(chan int)
	fmt.Println(getType(unknown))
	unknown = nil
	fmt.Println(getType(unknown))

}

func getType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
