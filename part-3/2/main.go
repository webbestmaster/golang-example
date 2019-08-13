package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func extractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))

	saiyans[1] = &Saiyan{
		Name:  "Say3",
		Power: 500,
	}

	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}

	return powers
}

func main() {
	sayanList := []*Saiyan{}

	sayanList = append(sayanList, &Saiyan{
		Name:  "Say1",
		Power: 100,
	})

	sayanList = append(sayanList, &Saiyan{
		Name:  "Say2",
		Power: 200,
	})

	fmt.Println(extractPowers(sayanList))

	fmt.Println(sayanList)
	fmt.Println(sayanList)
	fmt.Println(sayanList[0])
	fmt.Println(sayanList[1])
}
