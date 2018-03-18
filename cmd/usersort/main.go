package main

import (
	"os"

	"github.com/jsteenb2/anami"
)

func main() {
	s := anami.NewStudents([]string{"Dwayne", "David", "Jules", "Dana", "Riley", "Mohammad", "Kim", "Atlas", "Khadijah"})
	s.AddSubjects([]string{"English", "Math", "Science"})

	grades := map[string][]int{
		"Dwayne":   []int{70, 15, 45},
		"David":    []int{85, 50, 65},
		"Jules":    []int{74, 96, 88},
		"Dana":     []int{75, 90, 87},
		"Khadijah": []int{95, 100, 93},
		"Riley":    []int{69, 70, 71},
		"Mohammad": []int{90, 98, 92},
		"Kim":      []int{87, 80, 75},
		"Atlas":    []int{12, 31, 44},
	}
	s.AddGrades(grades)

	s.RunSort(os.Stdout)
}
