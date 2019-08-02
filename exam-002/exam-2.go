package main

import "fmt"

func main() {
	var a1 [3]int8
	fmt.Println(a1)

	a2 := [...]int8{1, 2, 3}
	fmt.Println(a2)

	var aa [3][2]int8
	aa[1][1] = 11
	fmt.Println(aa)

	var s1 []int
	s1 = append(s1, 11, 23, 23)
	fmt.Println(s1, len(s1), cap(s1))

	s2 := []int{10, 20, 30}
	fmt.Println(s2, len(s2), cap(s2))
	s2 = append(s2, 41, 43, 43)
	fmt.Println(s2, len(s2), cap(s2))

	sl3 := append(s1, s2...)
	fmt.Println(s2, len(sl3), cap(sl3))

	sl4 := make([]int, 7)
	fmt.Println(sl4, len(sl4), cap(sl4))

	sl5 := make([]int, len(sl4), len(sl4))
	copy(sl5, sl4)
	fmt.Println(sl5, len(sl5), cap(sl5))

	fmt.Println("Hello-2")
}
