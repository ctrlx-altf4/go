package main

import "fmt"

type Car struct {
	Name    string
	Brand   string
	CarType string // a field to separate
}

func (c *Car) GetName() string {
	return c.Name
}
func (c *Car) GetBrand() string {
	return c.Brand
}

type Bike struct {
	Name     string
	Brand    string
	BikeType string
}

type Details interface {
	GetName() string
	GetBrand() string
}

func (b *Bike) GetName() string {
	return b.Name
}

func (b *Bike) GetBrand() string {
	return b.Brand
}

func getDetails(v Details) string {
	return fmt.Sprintf(`Name: %s \n Brand: %s`, v.GetName(), v.GetBrand())
}

func main() {
	bike := Bike{Name: "Pulsar", Brand: "Yamaha", BikeType: "khai Bike ho"}
	car := Car{Name: "Car Name", Brand: "Suzuki", CarType: "This is a Car"}

	fmt.Println(bike)
	fmt.Println(car)

}
