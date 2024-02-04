package solutions

import "fmt"

type BigTruck struct{}

func (bt *BigTruck) LoadCargo(cargo int) string {
	return fmt.Sprintf("Грузовик загрузил %d единиц груза", cargo)
}

type SmallCar struct{}

func (sc *SmallCar) LoadCargo(cargo int) string {
	if cargo > 5 {
		return "Машина не может загрузить такой большой груз"
	}
	return fmt.Sprintf("Машина загрузила %d единиц груза", cargo)
}

type CarTrailer struct{}

func (t *CarTrailer) LoadCargo(cargo int) string {
	return fmt.Sprintf("Грузовик загрузил %d единиц груза", cargo)
}

type CarAdapter struct {
	car     SmallCar
	trailer CarTrailer
}

func (ca *CarAdapter) LoadCargo(cargo int) string {
	if cargo <= 5 {
		return ca.car.LoadCargo(cargo)
	} else {
		return fmt.Sprintf("%s, используя прицеп. Загружено %d единиц груза",
			ca.car.LoadCargo(5), cargo-5) + " от прицепа"
	}
}

func Example21() {
	truck := &BigTruck{}
	car := &SmallCar{}
	trailer := &CarTrailer{}
	carWithTrailer := &CarAdapter{car: *car, trailer: *trailer}

	fmt.Println(truck.LoadCargo(10))
	fmt.Println(car.LoadCargo(7))
	fmt.Println(carWithTrailer.LoadCargo(10))
}
