package main

import "fmt"

func main() {
	fmt.Println("Hello-4")
	fmt.Println(len("привет мир!"))

	i := 5
	b := &i

	*b = 6

	fmt.Println(i)

	sl := []int{0, 0, 0}
	ChangeSlice(sl)
	fmt.Println(sl)
}

// ChangeSlice change slice
func ChangeSlice(items []int) {
	items[2] = 1
}
