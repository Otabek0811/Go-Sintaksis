package main

import "fmt"

type Sweet struct {
	sugar, weight float64
	chocolate     bool
}

type Cake struct {
	Sweet
	name  string
	shape string
}

type chocolate struct {
	Sweet
	name  string
	cocos bool
}

func (s Sweet) getSugar() string {
	return fmt.Sprintf("Sugar: %fgr", s.sugar)
}

func (s Sweet) getWeight() string {
	return fmt.Sprintf("Sugar: %fgr", s.weight)
}

func main() {
	var Medoviy = Cake{
		Sweet: Sweet{
			sugar:     40,
			weight:    700,
			chocolate: false,
		},
		name:  "Medovik",
		shape: "circle",
	}

	var bounty = chocolate{
		Sweet: Sweet{
			sugar:     30,
			weight:    600,
			chocolate: true,
		},
		name:  "Bounty",
		cocos: true,
	}

	fmt.Println(Medoviy.getSugar())
	fmt.Println(bounty.getWeight())
}
