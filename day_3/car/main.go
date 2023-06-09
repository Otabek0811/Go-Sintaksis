package main

import "fmt"

type Car struct {
	Fuel          int
	milage_fuel   int
	milage_km     int
	oil_change    int
	oil_change_km int
}

func (c Car) expend(x int) (int, int) {
	if c.milage_km-c.oil_change_km>c.oil_change {
		fmt.Println("Change Oil!!!")
		return c.milage_fuel, c.Fuel 
	}
	
	return c.milage_fuel * x, c.Fuel - x
}

func (cr Car) oilchange(x int) int{
	return x

}

func main() {

	var bmw Car = Car{
		Fuel:          200,
		milage_fuel:   12,
		milage_km:     0,
		oil_change:    500,
		oil_change_km: 0,
	}
	km, Fuel := bmw.expend(55)
	bmw.milage_km+=km
	bmw.Fuel=Fuel
	bmw.oil_change_km+=bmw.oilchange(bmw.milage_km)
	km, Fuel = bmw.expend(55)
	bmw.milage_km+=km
	bmw.Fuel=Fuel
	bmw.oil_change_km+=bmw.oilchange(bmw.milage_km)
	// km, Fuel = bmw.expend(5)
	// bmw.milage_km+=km
	// bmw.Fuel=Fuel

	fmt.Println("Fuel: ",bmw.Fuel,"\tMilage km:",bmw.milage_km,"\tkm of change oil:",bmw.oil_change_km)
}
