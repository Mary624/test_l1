package solutions

import (
	"fmt"
	"reflect"
)

func Example14() {
	var ch chan string
	// с помощью fmt
	fmt.Println(typeOf1(ch))
	// с помощью type switch
	fmt.Println(typeOf2(ch))
	// с помощью reflect
	fmt.Println(typeOf3(ch))
}

func typeOf1(v any) string {
	return fmt.Sprintf("%T", v)
}

func typeOf2(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	}
	return "unknown"
}

func typeOf3(v any) string {
	return reflect.TypeOf(v).String()
}
