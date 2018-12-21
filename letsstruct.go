package main

import (
	"fmt"
)

func main() {

	type Saiyan struct {
		Name  string
		Power int
	}

	goku := &Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
	s.Power += 10000
}
