// genpass - strong password package
package genpass

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// func GetStrongPass returns as a string a password of random symbols,
// lowercase and uppercase letters with the length of (n int) and error,
// max value for n is 32767
func GetPass(n int) (string, error) {
	secNew := time.Now().UnixNano()
	rand.Seed(secNew)
	pass := ""
	if n > math.MaxInt16 {
		return "", fmt.Errorf("error: number is so big")
	}
	for i := 0; i < n; i++ {
		r := rand.Intn(4)
		switch r {
		case 0:
			pass += letterUp()
		case 1:
			pass += digt()
		case 2:
			pass += letterLow()
		case 3:
			pass += simbol()
		}
	}
	return pass, nil
}

func digt() string {
	num := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	x := random(len(num))
	return num[x]
}

func simbol() string {
	sim := []string{"!", "@", "#", "$", "%", "&", "?"}
	x := random(len(sim))
	return sim[x]
}

func letterUp() string {
	lettU := []string{"A", "B", "C", "D", "E", "F", "G", "H",
		"I", "J", "K", "L", "M", "N", "O", "P", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	x := random(len(lettU))
	return lettU[x]
}

func letterLow() string {
	lettL := []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	x := random(len(lettL))
	return lettL[x]
}

func random(lenStr int) int {
	seconds := time.Now().UnixNano()
	rand.Seed(seconds)
	r := rand.Intn(lenStr)
	return r
}