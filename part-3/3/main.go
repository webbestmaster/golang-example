package main

import "fmt"

func main() {
	/*
	   lookup := map[string]int{
	   	"goku": 9001,
	   	"gohan": 2044,
	   }
	*/
	lookup := make(map[string]int, 1)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	// power == 0, exists == false
	// 0 это значение по умолчанию для типа integer
	fmt.Println(power, exists)
}
