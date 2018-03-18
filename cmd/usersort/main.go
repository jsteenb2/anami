package main

import (
	"os"

	"github.com/jsteenb2/anami"
)

func main() {
	s := anami.NewStudents([]string{"Dwayne", "David", "Jules", "Dana", "Riley", "Mohammad", "Kim", "Atlas", "Khadijah"})
	s.AddSubjects([]string{"English", "Math", "Science", "History", "Arabic", "Gym"})

	grades := map[string][]int{
		"Dwayne":   []int{70, 15, 45, 73, 71, 91},
		"David":    []int{85, 50, 65, 81, 56, 81},
		"Jules":    []int{74, 96, 88, 90, 37, 83},
		"Dana":     []int{75, 90, 87, 77, 77, 90},
		"Khadijah": []int{95, 100, 93, 99, 100, 97},
		"Riley":    []int{69, 70, 71, 78, 83, 85},
		"Mohammad": []int{90, 98, 92, 96, 100, 99},
		"Kim":      []int{87, 80, 75, 66, 91, 73},
		"Atlas":    []int{12, 31, 44, 80, 80, 31},
	}
	s.AddGrades(grades)

	s.RunSort(os.Stdout)
}
