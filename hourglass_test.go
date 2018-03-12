package anami_test

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/jsteenb2/anami"
)

var update = flag.Bool("update", false, "update .golden files")

func Test_NewHourGlass(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		fileName string
	}{
		{"7 rows", 7, "7_rows.golden"},
		{"5 rows", 5, "5_rows.golden"},
		{"9 rows", 9, "9_rows.golden"},
		{"31 rows", 31, "31_rows.golden"},
		{"6 rows", 6, "6_rows.golden"},
		{"20 rows", 20, "20_rows.golden"},
		{"50 rows", 50, "50_rows.golden"},
		{"300 rows", 300, "300_rows.golden"},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			actual := anami.NewHourGlass(tt.in)

			expected := goldenHelper(t, filepath.Join("golden", tt.fileName), []byte(actual))
			if actual != string(expected) {
				t.Errorf("\ngot:\n%s\nexpected:\n%s", actual, expected)
			}
		}
		t.Run(tt.name, fn)
	}
}

func Test_NewAnamiRow(t *testing.T) {
	tests := []struct {
		name               string
		lineLen, totalRows int
		expected           string
	}{
		{"1st row of 7 rows", 7, 7, "ANAMIAN"},
		{"2nd row of 7 rows", 5, 7, " ANAMI "},
		{"3rd row of 7 rows", 3, 7, "  ANA  "},
		{"4th row of 7 rows", 1, 7, "   A   "},
		{"5th row of 7 rows", 3, 7, "  ANA  "},
		{"6th row of 7 rows", 5, 7, " ANAMI "},
		{"7th row of 7 rows", 7, 7, "ANAMIAN"},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			if s := anami.NewAnamiRow(tt.lineLen, tt.totalRows); s != tt.expected {
				t.Errorf("\ngot: %s\nwanted: %s", s, tt.expected)
			}
		}
		t.Run(tt.name, fn)
	}
}

func Test_NextPrime(t *testing.T) {
	tests := []struct {
		name    string
		in, out int
	}{
		{"simple case", 2, 3},
		{"edge case of 1", 1, 2},
		{"edge case of 0", 0, 2},
		{"from 2", 2, 3},
		{"from 3", 3, 5},
		{"from 5", 5, 7},
		{"from 7", 7, 11},
		{"from 90", 90, 97},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			if p := anami.NextPrime(tt.in); p != tt.out {
				t.Errorf("\ngot: %d\nwanted: %d", p, tt.out)
			}
		}
		t.Run(tt.name, fn)
	}
}

func Test_NewPrimeRow(t *testing.T) {
	tests := []struct {
		name                     string
		lineLen, totalRows       int
		prevPrime, expectedPrime int
		carry, expectedCarry     string
		expectedLine             string
	}{
		{"1st row of 7 rows", 7, 7, 0, 7, "", "", "2+3+5+7"},
		{"2nd row of 7 rows", 5, 7, 7, 13, "", "", " 11+13 "},
		{"3rd row of 7 rows", 3, 7, 13, 17, "", "", "  17+  "},
		{"4th row of 7 rows", 1, 7, 17, 19, "", "9", "   1   "},
		{"5th row of 7 rows", 3, 7, 19, 23, "9", "3", "  9+2  "},
		{"6th row of 7 rows", 5, 7, 23, 29, "3", "", " 3+29+ "},
		{"7th row of 7 rows", 7, 7, 29, 41, "", "1", "31+37+4"},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			s, carry, p := anami.NewPRow(tt.lineLen, tt.totalRows, tt.prevPrime, tt.carry)
			if s != tt.expectedLine {
				t.Errorf("\ngot line: %s\nwanted: %s", s, tt.expectedLine)
			}

			if carry != tt.expectedCarry {
				t.Errorf("\ngot carry: %s\nwanted: %s", carry, tt.expectedCarry)
			}

			if p != tt.expectedPrime {
				t.Errorf("\ngot prime: %d\nwanted: %d", p, tt.expectedPrime)
			}
		}
		t.Run(tt.name, fn)
	}
}
