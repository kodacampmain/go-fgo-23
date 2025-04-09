package main

import (
	"errors"
	"fmt"
	"time"
)

// var globals string

func main() {
	// fl := true
	// uid := uuid.NewString()
	// uid = ""
	// Print(uid)
	// a.PrintAnimal()
	// var firstname string
	// fmt.Println(firstname)
	// firstname = "Fazz"
	// fmt.Println(firstname)
	// var lastname string = "Track"
	// fmt.Println(lastname)
	// var age uint8
	// fmt.Println(age)
	// age = 20
	// fmt.Println(age)
	// fmt.Println(globals)
	// isMarried := false
	// if isMarried {
	// 	fmt.Println("Sudah Menikah")
	// } else {
	// 	fmt.Println("Belum Menikah")
	// }
	// if age > 12 && age < 17 {
	// 	fmt.Println("Remaja")
	// } else if age > 17 {
	// 	fmt.Println("Dewasa")
	// } else {
	// 	fmt.Println("Anak-anak")
	// }
	// varian for loop
	// for i := 1; i <= 5; i++ {
	// 	var row string
	// 	// varian while loop
	// 	j := 0
	// 	for j < i {
	// 		row += "*"
	// 		j++
	// 	}
	// 	// END varian while loop
	// 	fmt.Println(row)
	// }
	// // END varian for loop
	// for i := 0; ; {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	fmt.Print(i)
	// 	i++
	// }
	// var hobbies [3]string

	// hobbies[0] = "Swimming"
	// hobbies[2] = "Watching"

	// fmt.Printf(hobbies[1])
	// var scores [10]uint8
	// fmt.Println(scores)
	// fmt.Println(len(scores))

	// fruits := [3]string{"Semangka", "Melon", "Durian"}
	// for idx, fruit := range fruits {
	// 	fmt.Printf("%d. %s\n", idx+1, fruit)
	// }
	// favFruits := fruits[:]
	// fruits[2] = "Pisang"
	// favFruits = append(favFruits, "Durian")
	// favFruits = slices.Delete(favFruits, 1, 2)
	// // oldLen := len(favFruits)
	// // favFruits = append(favFruits[:1], favFruits[2:]...)
	// // clear(favFruits[len(favFruits):oldLen])
	// newFruits := []string{"Mangga", "Jambu"}
	// favFruits = slices.Replace(favFruits, 2, 2, newFruits...)
	// // favFruits = append(favFruits[:2], append(newFruits, favFruits[2:]...)...)
	// fmt.Println(favFruits)
	// // fmt.Println(fruits)

	// // copy()

	// var empty []int
	// fmt.Println(empty)

	var person = map[string]string{}
	person["name"] = "Iman"
	person["umur"] = "40"

	fmt.Println(person)

	person2 := map[string]string{
		"name": "Sandi",
		"umur": "5",
	}
	fmt.Println(person2)

	person3 := make(map[string]string)
	person3["name"] = "Hakim"
	person3["umur"] = "120"
	fmt.Println(person3)
	printSegitiga(10)
	result, err := getMovies()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func printSegitiga(s int) {
	// varian for loop
	for i := 1; i <= s; i++ {
		var row string
		// varian while loop
		j := 0
		for j < i {
			row += "*"
			j++
		}
		// END varian while loop
		fmt.Println(row)
	}
	// END varian for loop
}

func getMovies() ([]string, error) {
	movies := []string{"Minecraft", "Frozen", "John Wick"}
	isSuccess := true
	time.Sleep(1000 * time.Millisecond)
	if isSuccess {
		return movies, nil
	}
	return nil, errors.New("internal server error")
}
