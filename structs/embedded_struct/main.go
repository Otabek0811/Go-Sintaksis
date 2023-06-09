package main

import "fmt"

type Component struct{
	weight, sugar float64
}

type Fruit struct{
	Component
	name string
}

type Vegeatable struct{
	Component
	name string
}

func (c Component) getSugar()string{
	return fmt.Sprintf("Sugar: %fgr",c.sugar)
}




func main() {
	var apple = Fruit{
		Component: Component{100,19},
		name: "Apple",
	}

	var tomato = Vegeatable{
		Component: Component{100,2.94},
		name: "Tomato",
	}

	fmt.Println(apple.getSugar())
	fmt.Println(tomato.getSugar())
}
