package models

import "fmt"

type Cat struct {
	Race    string
	Weight  uint
	Pattern string
}

func (c Cat) Attack() {
	fmt.Println("Cat attacked you for 7 damage")
}

func (c Cat) Agility() {
	fmt.Println("Cat uses agility, evasion increases")
}

func (c Cat) Sound() {
	fmt.Println("RAAAAAWR")
}
