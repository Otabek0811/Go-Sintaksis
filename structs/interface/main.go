package main

import "fmt"

type Os struct {
	Name string
}

type Linux struct {
	Os
}

type MacOs struct {
	Os
}

type Windows struct {
	Os
}

type Osinterface interface {
	GetName() string
}

func (o Os) GetName() string {
	return o.Name
}

func (o MacOs) GetName() string {
	// Super call
	return "Macbook" + o.Os.GetName()
}

func Run(os Osinterface) {
	fmt.Println("Loading " + os.GetName() + ".......")

}

func main() {
	Run(Linux{Os{Name: "Ubuntu 22.04"}})
	Run(MacOs{Os{Name: " Pro M1"}})
	Run(Windows{Os{Name: "Windows 10 Pro"}})
}
