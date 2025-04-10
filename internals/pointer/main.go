package pointer

import "fmt"

func ShowPointer() {
	var value int
	var point *int

	value = 10
	fmt.Printf("nilai awal: %d\n", value)
	point = &value
	pointB := &value
	fmt.Printf("reference value: %p\n", &value)
	fmt.Printf("nilai variabel point: %p\n", point)
	*point = 20 // jika mengubah nilai yg ditunjuk pointer, maka akan mengubah nilai asli yg bersangkutan
	valueB := value
	fmt.Printf("dereference point: %d\n", *point)
	fmt.Printf("dereference point B: %d\n", *pointB)
	fmt.Printf("nilai variabel value: %d\n", value)
	fmt.Printf("nilai variabel value B: %d\n", valueB)

}
