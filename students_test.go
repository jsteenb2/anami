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

func Test_Students_PrintDefault_SortsCorrectly(t *testing.T) {
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

func Test_Students_Print_SortsCorrectly(t *testing.T) {
	tests := []struct {
		name            string
		names, subjects []string
		grades          map[string][]int
		fileName        string
		subject         int
		sortOrder       anami.SortOrder
	}{
		{
			"Sort on 0th subject in ASC order",
			[]string{"Dwayne", "David", "Khadijah", "Jules"},
			[]string{"English", "Math", "Science"},
			map[string][]int{
				"Dwayne":   []int{70, 40, 45},
				"David":    []int{90, 50, 65},
				"Khadijah": []int{95, 100, 93},
				"Jules":    []int{74, 96, 88},
			},
			"sort_0thsubj_ASC.txt",
			0,
			anami.ASC,
		},
		{
			"Sort on 2nd subject in DESC order",
			[]string{"Dwayne", "David", "Jules", "Dana", "Riley", "Mohammad", "Kim", "Atlas", "Khadijah"},
			[]string{"English", "Math", "Science"},
			map[string][]int{
				"Dwayne":   []int{70, 15, 45},
				"David":    []int{85, 50, 65},
				"Jules":    []int{74, 96, 88},
				"Dana":     []int{75, 90, 87},
				"Khadijah": []int{95, 100, 93},
				"Riley":    []int{69, 70, 71},
				"Mohammad": []int{90, 98, 92},
				"Kim":      []int{87, 80, 75},
				"Atlas":    []int{12, 31, 44},
			},
			"sort_2ndsubj_DESC.txt",
			2,
			anami.DESC,
		},
	}

	sorts := []anami.SortType{anami.QuickSort, anami.MergeSort, anami.SelectionSort}

	for _, tt := range tests {
		for _, sortType := range sorts {
			fn := func(t *testing.T) {
				students := anami.NewStudents(tt.names)
				students.AddSubjects(tt.subjects)
				students.AddGrades(tt.grades)

				var b bytes.Buffer
				students.Print(&b, tt.sortOrder, tt.subject, sortType)

				expected := goldenHelper(t, filepath.Join("golden", tt.fileName), b.Bytes())
				if !bytes.Equal(b.Bytes(), expected) {
					t.Errorf("\ngot:\n%s\nexpected:\n%s", string(b.Bytes()), string(expected))
				}
			}
			t.Run(tt.name+" "+sortType.String(), fn)

			// only run the quicksort test with update
			if *update {
				break
			}
		}
	}
}
