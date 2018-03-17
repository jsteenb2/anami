package anami_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/jsteenb2/anami"
)

func Test_Students_AddGrades(t *testing.T) {
	tests := []struct {
		name            string
		names, subjects []string
		grades          map[string][]int
		fileName        string
	}{
		{
			"2 students",
			[]string{"Dwayne", "David"},
			[]string{"English", "Math", "Science"},
			map[string][]int{
				"Dwayne": []int{90, 70, 85},
				"David":  []int{74, 96, 88},
			},
			"grades_1.txt",
		},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			students := anami.NewStudents(tt.names)
			students.AddSubjects(tt.subjects)
			students.AddGrades(tt.grades)

			var b bytes.Buffer
			students.PrintDefault(&b)

			expected := goldenHelper(t, filepath.Join("golden", tt.fileName), b.Bytes())
			if !bytes.Equal(b.Bytes(), expected) {
				t.Errorf("\ngot:\n%s\nexpected:\n%s", string(b.Bytes()), string(expected))
			}
		}
		t.Run(tt.name, fn)
	}
}

func Test_Students_SortsCorrectly(t *testing.T) {
	tests := []struct {
		name            string
		names, subjects []string
		grades          map[string][]int
		fileName        string
	}{
		{
			"4 students",
			[]string{"Dwayne", "David", "Khadijah", "Jules"},
			[]string{"English", "Math", "Science"},
			map[string][]int{
				"Dwayne":   []int{70, 40, 45},
				"David":    []int{90, 50, 65},
				"Khadijah": []int{95, 100, 93},
				"Jules":    []int{74, 96, 88},
			},
			"grades_sorted_4_students.txt",
		},
		{
			"9 students",
			[]string{"Dwayne", "David", "Jules", "Dana", "Riley", "Mohammad", "Kim", "Atlas", "Khadijah"},
			[]string{"English", "Math", "Science"},
			map[string][]int{
				"Dwayne":   []int{70, 40, 45},
				"David":    []int{85, 50, 65},
				"Jules":    []int{74, 96, 88},
				"Dana":     []int{75, 90, 88},
				"Khadijah": []int{95, 100, 93},
				"Riley":    []int{70, 70, 71},
				"Mohammad": []int{90, 98, 92},
				"Kim":      []int{90, 80, 75},
				"Atlas":    []int{12, 31, 44},
			},
			"grades_sorted_9_students.txt",
		},
	}

	for _, tt := range tests {
		fn := func(t *testing.T) {
			students := anami.NewStudents(tt.names)
			students.AddSubjects(tt.subjects)
			students.AddGrades(tt.grades)

			var b bytes.Buffer
			students.PrintDefault(&b)

			expected := goldenHelper(t, filepath.Join("golden", tt.fileName), b.Bytes())
			if !bytes.Equal(b.Bytes(), expected) {
				t.Errorf("\ngot:\n%s\nexpected:\n%s", string(b.Bytes()), string(expected))
			}
		}
		t.Run(tt.name, fn)
	}
}
