package models

import "fmt"

type Car struct {
	Model  string
	Year   uint
	Color  string
	Engine uint
	Brand  string
}

func (c Car) GetCarFullName() string {
	carFullName := fmt.Sprintf("%s %s", c.Brand, c.Model)
	return carFullName
}

func (c Car) Sound() {
	fmt.Println("Telolet")
}

func (c Car) Attack() {
	fmt.Println("Car hits you for 100 damage")
}
