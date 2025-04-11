package minitask

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// var nama string
// var fotoUri string
// var fotoFile []byte
// var email string
// var umur uint8
// var telp string
// var isMarried bool
// var schools = []map[string]string{
// 	{"nama sekolah": "SMA", "jurusan": "IPA"},
// 	{"nama sekolah": "Universitas", "jurusan": "Metalurgi"},
// }

func FakeLogin(email, password string) (bool, error) {
	// userEmail := "mail@mail.com"
	// userPass := "abcdefgh"
	if email != adminEmail {
		return false, errors.New("email/password salah")
	}
	if password != adminPassword {
		return false, errors.New("email/password salah")
	}
	return true, nil
}

func SliceOfArr() {
	var arr = [5]uint{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Printf("array: %v\nslice: %v\n", arr, slice)
}

func MaxValueArray(arr [5]uint) uint {
	var max uint
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func TempConversion(temp float32, srcUnit string, destUnit string) (float32, error) {
	var result float32
	rgx, err := regexp.Compile("[KkFfRrCc]{1}")
	if err != nil {
		return 0, err
	}

	isMatched := rgx.Match([]byte(srcUnit))
	if !isMatched {
		return 0, errors.New("unit tidak sesuai")
	}

	isMatched = rgx.Match([]byte(destUnit))
	if !isMatched {
		return 0, errors.New("unit tidak sesuai")
	}

	if strings.EqualFold(srcUnit, destUnit) {
		return temp, nil
	}

	var mappedUnit = map[string]float32{
		"K": 5,
		"C": 5,
		"R": 4,
		"F": 9,
	}

	if srcUnit == "K" {
		result = temp - 273
	} else if srcUnit == "F" {
		result = temp - 32
	} else {
		result = temp
	}

	result = result * mappedUnit[destUnit] / mappedUnit[srcUnit]

	if destUnit == "K" {
		result += 273
	} else if destUnit == "F" {
		result += 32
	}

	return result, nil
}
