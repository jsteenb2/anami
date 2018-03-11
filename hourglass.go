package anami

import (
	"fmt"
	"strings"
)

var aNAMI = "ANAMI"

func NewHourGlass(size int) string {
	var (
		output, carry string
		p             int
	)
	for row := 0; row < size; row++ {
		lineLen := getLineLen(row, size)
		if lineLen == 0 {
			continue
		}
		var line string
		line, carry, p = NewPRow(lineLen, size, p, carry)
		output += line + " " + NewAnamiRow(lineLen, size) + "\n"
	}

	return output
}

func NewAnamiRow(lineLen, total int) string {
	return padLine(aline(lineLen).String(), lineLen, total)
}

type aline int

func (a aline) String() string {
	s := strings.Repeat(aNAMI, int(a)/len(aNAMI))
	s += aNAMI[:int(a)%len(aNAMI)]
	return s
}

func NewPRow(lineLen, total, p int, s string) (line, carry string, lastPrime int) {
	l, c, p := newPLine(lineLen, p, s)
	return padLine(l, lineLen, total), c, p
}

func newPLine(lineLen, p int, s string) (line, carry string, lastPrime int) {
	if len(s) != 0 {
		s += "+"
	}

	for len(s) < lineLen {
		p = NextPrime(p)
		s += fmt.Sprintf("%d+", p)
	}

	if len(s) > lineLen {
		return s[:lineLen], strings.Replace(s[lineLen:], "+", "", -1), p
	}
	return s, "", p
}

func NextPrime(i int) int {
	if i < 2 {
		return 2
	}

	var newPrime int
NewPrimeLoop:
	for k := i + 1; ; k++ {
		for j := 2; j < k; j++ {
			if k%j == 0 {
				continue NewPrimeLoop
			}
		}
		newPrime = k
		break NewPrimeLoop
	}

	return newPrime
}

func getLineLen(row, size int) int {
	if row < size/2 {
		return size - row*2
	}

	if size % 2 == 0 {
		return 2 + 2*(row-size/2)
	}

	return 1 + 2*(row-size/2)
}

func padLine(s string, lineLen, total int) string {
	return strings.Repeat(" ", (total-lineLen)/2) + s + strings.Repeat(" ", (total-lineLen)/2)
}
