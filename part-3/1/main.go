package main

import "fmt"

func main() {
	var scores [10]int
	scores[4] = 339

	fmt.Println(scores)

	for index, value := range scores {
		fmt.Println("index: ", index, " value: ", value)
	}

	scores2 := []int{1, 4, 293, 4, 9}

	for index, value := range scores2 {
		fmt.Println("index: ", index, " value: ", value)
	}
}
