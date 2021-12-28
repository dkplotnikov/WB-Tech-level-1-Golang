package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func myTypeOf(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown type"
	}
}

func main() {

	X := []interface{}{1, "hello", true, make(chan int)}

	for _, x := range X {
		fmt.Println("myTypeOf:\t", myTypeOf(x))
	}

	for _, x := range X {
		fmt.Println("reflect.TypeOf:\t", reflect.TypeOf(x))
	}

	for _, x := range X {
		fmt.Println("reflect.ValueOf.Kind:\t", reflect.ValueOf(x).Kind())
	}

	for _, x := range X {
		fmt.Printf("Printf %%T:\t%T\n", x)
	}

}
