package main

import (
	"fmt"
	"reflect"
)

type Slice struct {
	X string
	Y string
}

func main() {
	s := Slice{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		value := reflect.ValueOf(field)
		value.SetString("adsfasf")
		fmt.Println(value.String())
	}
}
