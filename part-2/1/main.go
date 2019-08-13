package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)

	goku2 := NewSaiyan("Hocku", 2000)
	Super(&goku2)
	fmt.Println(goku2.Power)
}

func Super(s *Saiyan) {
	s.Power += 10000
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
