package main

import (
	"fgo23/internals/races"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// var globals string

func main() {
	defer fmt.Println("Program Selesai")
	fmt.Println("Program Mulai")
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

	// var person = map[string]string{}
	// person["name"] = "Iman"
	// person["umur"] = "40"

	// fmt.Println(person)

	// person2 := map[string]string{
	// 	"name": "Sandi",
	// 	"umur": "5",
	// }
	// fmt.Println(person2)

	// person3 := make(map[string]string)
	// person3["name"] = "Hakim"
	// person3["umur"] = "120"
	// fmt.Println(person3)
	// printSegitiga(10)
	// printSegitigaRecursive(1, 10)
	// result, err := getMovies()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println(printFibo(30))
	// fmt.Println(printFiboRecursive([]uint{0, 1}, 30))
	// pointer.ShowPointer()
	// srcTemp := float32(300)
	// srcUnit := "K"
	// destUnit := "R"
	// tempResult, err := minitask.TempConversion(srcTemp, srcUnit, destUnit)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("Konversi suhu %f dari %s ke %s adalah %f\n", srcTemp, srcUnit, destUnit, tempResult)
	// }
	// fmt.Printf("hasil %d", minitask.MaxValueArray([5]uint{10, 5, 17, 90, 69}))
	// var spiderman models.Movie
	// spiderman := models.NewMovie("Spiderman", "spiderman.jpg", []string{"Tom Holland", "Zendaya"}, []string{"Action", "Romance", "Drama"}, 120, 50000)
	// fmt.Println(spiderman.GetMovieTitle())
	// fmt.Println(spiderman.GetMoviePrice())
	// spiderman.ChangePrice(20000)
	// fmt.Println(spiderman.GetMoviePrice())

	// printSound(models.Car{})
	// printSound(models.Cat{})

	// var score = [5]int{70, 80, 65, 90, 100}

	// fmt.Println(minitask.SumArray(score))
	// fmt.Println(minitask.SumArrayRecursive(score, 0))
	// age := 10
	// fmt.Println(age)
	// minitask.ChangeNum(&age, 17)
	// fmt.Println(age)
	// man := &minitask.Person{
	// 	Name:    "",
	// 	Address: "",
	// 	Phone:   "",
	// }

	// manusia := minitask.NewPerson("Fazztrack", "Legenda Wisata", "08123456789")
	// fmt.Println(manusia.PrintPersonInfo())
	// manusia.GreetPerson()
	// manusia.EditName("Koda")
	// manusia.GreetPerson()

	// conStr := minitask.GeneratePgConnString(minitask.DBConfig{User: "fakhridho", DbName: "movies"})
	// fmt.Println(conStr)

	// var wg sync.WaitGroup

	// Add => menambah antrian
	// Done => mengurangi antrian
	// Wait => menunggu hingga antrian kosong

	// Channel => saluran komunikasi antar goroutine
	// make()
	// manifest type => chan datatype

	// var numChannel = make(chan int, 5)

	// go PrintRandomNumber(5, 6, numChannel)

	// urls := []string{
	// 	"https://stackoverflow.com",
	// 	"https://github.com",
	// 	"https://www.linkedin.com",
	// 	"https://medium.com",
	// 	"https://golang.org",
	// 	"https://www.fazztrack.com",
	// 	"https://www.coursera.org",
	// 	"https://wesionary.team",
	// }
	// strChannel := make(chan string, len(urls))

	// wg.Add(1)
	// go PingWebsite(urls, strChannel)
	// fmt.Println(<-strChannel)

	// for {
	// 	if _, ok := <-numChannel; !ok {
	// 		break
	// 	}
	// }
	// for {
	// 	if _, ok := <-strChannel; !ok {
	// 		break
	// 	}
	// }

	// closedChannel := &channels{
	// 	numChannel: true,
	// 	strChannel: true,
	// }

	// for {
	// 	select {
	// 	case _, ok := <-numChannel:
	// 		if !ok {
	// 			fmt.Println("numChannel out")
	// 			closedChannel.closeNum()
	// 		}
	// 	case _, ok := <-strChannel:
	// 		if !ok {
	// 			fmt.Println("strChannel out")
	// 			closedChannel.closeStr()
	// 		}
	// 	default:
	// 	}
	// 	if !closedChannel.numChannel && !closedChannel.strChannel {
	// 		break
	// 	}
	// }
	// for num := range numChannel {
	// 	fmt.Println(num)
	// }
	// for str := range strChannel {
	// 	fmt.Println(str)
	// }

	// wg.Wait()
	// time.Sleep(2000 * time.Millisecond)

	races.Run()

}

func PrintRandomNumber(count int, limit int, chn chan int) {
	// defer wg.Done()
	i := 1
	for i <= count {
		result := rand.Intn(limit)
		chn <- result
		fmt.Println(result)
		i++
	}
	close(chn)
}

func PingWebsite(urls []string, chn chan string) {
	// defer wg.Done()
	var wg sync.WaitGroup
	defer close(chn)
	defer wg.Wait()
	// defer LIFO (last in first out)

	for _, url := range urls {
		// fmt.Println("Fetching ", url)
		// defer fmt.Println("Done Fetching", url)
		wg.Add(1)
		go func(url string, chn chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			start := time.Now()
			res, err := http.Get(url)
			if err != nil {
				str := fmt.Sprintf("Get %s mengalami Error, Started at %d\n", url, start.Second())
				chn <- str
				fmt.Print(str)
				return
			}
			str := fmt.Sprintf("[%d] %s is UP, Started at %d\n", res.StatusCode, url, start.Second())
			chn <- str
			fmt.Print(str)
		}(url, chn, &wg)
	}
}

// func printSegitiga(s int) {
// 	// varian for loop
// 	for i := 1; i <= s; i++ {
// 		var row string
// 		// varian while loop
// 		j := 0
// 		for j < i {
// 			row += "*"
// 			j++
// 		}
// 		// END varian while loop
// 		fmt.Println(row)
// 	}
// 	// END varian for loop
// }

// func printSegitigaRecursive(i, s int) {
// 	if i > s {
// 		return
// 	}
// 	row := printSegitigaRowRecursive(0, i)
// 	fmt.Println(row)
// 	printSegitigaRecursive(i+1, s)
// }

// func printSegitigaRowRecursive(j, i int) string {
// 	if j >= i {
// 		return ""
// 	}
// 	return "*" + printSegitigaRowRecursive(j+1, i)
// }

// func getMovies() ([]string, error) {
// 	movies := []string{"Minecraft", "Frozen", "John Wick"}
// 	isSuccess := true
// 	time.Sleep(1000 * time.Millisecond)
// 	if isSuccess {
// 		return movies, nil
// 	}
// 	return nil, errors.New("internal server error")
// }

// func printFibo(max uint) []uint {
// 	seq := []uint{0, 1}
// 	seqLen := len(seq)
// 	// index n
// 	// index n-1 + index n-2
// 	for {
// 		nextSeq := seq[seqLen-1] + seq[seqLen-2]
// 		if nextSeq > max {
// 			break
// 		}
// 		seq = append(seq, nextSeq)
// 		seqLen = len(seq)
// 	}
// 	return seq
// }

// func printFiboRecursive(seq []uint, max uint) []uint {
// 	seqLen := len(seq)
// 	nextSeq := seq[seqLen-1] + seq[seqLen-2]
// 	if nextSeq > max {
// 		return seq
// 	} else if nextSeq == max {
// 		return append(seq, nextSeq)
// 	}
// 	newSeq := append(seq, nextSeq)
// 	return printFiboRecursive(newSeq, max)
// }

// type sesuatuYangBersuara interface {
// 	Sound()
// 	Attack()
// }

// func printSound(t sesuatuYangBersuara) {
// 	t.Sound()
// }

// func printCatSound(t sesuatuYangBersuara) {
// 	t.Sound()
// }
