package minitask

import "fmt"

func SumArray(arr [5]int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumArrayRecursive(arr [5]int, i uint) int {
	if i >= uint(len(arr)) {
		return 0
	}
	return arr[i] + SumArrayRecursive(arr, i+1)
}

func ChangeNum(num *int, newNum int) {
	// untuk mengubah value dari pointer
	// pointer diubah dulu menjadi value menggunakan dereference
	*num = newNum
}

type Person struct {
	Name      string
	Address   string
	Phone     string
	Age       uint
	IsMarried bool
}

// constructor
func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) PrintPersonInfo() string {
	return fmt.Sprintf("Nama: %s\nAlamat: %s\nNo Telepon: %s\n", p.Name, p.Address, p.Phone)
}

func (p *Person) GreetPerson() {
	fmt.Printf("Selamat Pagi %s\n", p.Name)
}

func (p *Person) EditName(newName string) {
	p.Name = newName
}

type DBConfig struct {
	User   string
	Pass   string
	Host   string
	Port   uint
	DbName string
}

func GeneratePgConnString(config DBConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.User, config.Pass, config.Host, config.Port, config.DbName)
}

var adminEmail string = "admin@mail.com"
var adminPassword string = "admin123"
