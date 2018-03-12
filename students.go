package anami

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type students struct {
	subjects []string
	s        map[string]*student
}

func NewStudents(names []string) *students {
	s := make(map[string]*student)
	for _, name := range names {
		s[name] = &student{name: name, grades: make([]string, 0)}
	}

	return &students{
		subjects: make([]string, 0),
		s:        s,
	}
}

func (s *students) AddSubjects(subs []string) {
	s.subjects = append(s.subjects, subs...)
}

func (s *students) AddGrades(grades map[string][]int) {
	for name, gs := range grades {
		for _, v := range gs {
			s.s[name].grades = append(s.s[name].grades, strconv.Itoa(v))
			s.s[name].totalMarks += v
		}
	}
}

func (s *students) Print(w io.Writer) {
	table := newTable(w)
	table.SetHeader(s.headers())

	names := s.sortedStudentNames()
	for i, name := range names {
		row := []string{strconv.Itoa(i + 1), name}
		student := s.s[name]
		if len(student.grades) > 0 {
			row = append(row, student.grades...)
			row = append(row, strconv.Itoa(student.totalMarks))
		}
		table.Append(row)
	}

	table.Render()
}

func (s *students) headers() []string {
	h := append([]string{"Rank", "Name"}, s.subjects...)
	h = append(h, "Total Marks")
	return h
}

func (s *students) sortedStudentNames() []string {
	var names []string
	for k := range s.s {
		names = append(names, k)
	}

	return s.quicksort(names)
}

// sorts in descending order
func (s *students) quicksort(n []string) []string {
	if len(n) < 2 {
		return n
	}

	p := s.getPivot(n)
	pivStudent := s.s[n[p]]
	n = swap(n, p)

	var lastMoved int
	for i := 0; i < len(n)-1; i++ {
		curStudent := s.s[n[i]]
		if curStudent.totalMarks > pivStudent.totalMarks {
			n[lastMoved], n[i] = n[i], n[lastMoved]
			lastMoved++
		}
	}
	n = swap(n, lastMoved)
	out := append(s.quicksort(n[:lastMoved]), n[lastMoved])
	out = append(out, s.quicksort(n[lastMoved+1:])...)
	return out
}

func (s *students) getPivot(n []string) int {
	first := s.s[n[0]]
	mid := s.s[n[len(n)/2]]
	last := s.s[n[len(n)-1]]

	if first.totalMarks > mid.totalMarks && first.totalMarks < last.totalMarks {
		return 0
	} else if mid.totalMarks > last.totalMarks {
		return len(n) - 1
	}

	return len(n) / 2
}

func swap(n []string, i int) []string {
	n[len(n)-1], n[i] = n[i], n[len(n)-1]
	return n
}

type student struct {
	name       string
	grades     []string
	totalMarks int
}

func newTable(w io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(w)
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)
	table.SetCenterSeparator("-")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return table
}
