package anami

import (
	"io"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
)

const (
	ASC SortOrder = iota
	DESC

	QuickSort SortType = iota
	MergeSort
	SelectionSort
	HeapSort
)

type SortOrder int

func NewSortOrder(i int) SortOrder {
	switch i {
	case 0:
		return ASC
	case 1:
		return DESC
	default:
		return SortOrder(i)
	}
}

type SortType int

func (s SortType) String() string {
	var out string

	switch int(s) {
	case int(QuickSort):
		out = "Quick Sort"
	case int(MergeSort):
		out = "Merge Sort"
	case int(SelectionSort):
		out = "Selection Sort"
	case int(HeapSort):
		out = "Heap Sort"
	}

	return out
}

func NewSortType(i int) SortType {
	switch i {
	case 0:
		return QuickSort
	case 1:
		return MergeSort
	case 2:
		return SelectionSort
	case 3:
		return HeapSort
	default:
		return SortType(i)
	}
}

type students struct {
	subjects []string
	s        map[string]*student
}

func NewStudents(names []string) *students {
	s := make(map[string]*student)
	for _, name := range names {
		s[name] = &student{name: name, grades: make([]int, 0)}
	}

	return &students{
		subjects: make([]string, 0),
		s:        s,
	}
}

func (s *students) RunSort(w io.Writer) error {
	prompt := promptui.Select{
		Label: "Select Order",
		Items: []string{
			"ASC",
			"DESC",
		},
	}
	i, _, err := prompt.Run()
	if err != nil {
		return err
	}
	order := NewSortOrder(i)

	prompt = promptui.Select{
		Label: "Select Subject",
		Items: append(s.subjects, "Total Marks"),
	}
	subj, _, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Select{
		Label: "Select Sort Algorithm",
		Items: []string{QuickSort.String(), MergeSort.String(), SelectionSort.String(), HeapSort.String()},
	}
	k, _, err := prompt.Run()
	if err != nil {
		return err
	}
	sortType := NewSortType(k)

	s.Print(w, order, subj, sortType)
	return nil
}

func (s *students) AddSubjects(subs []string) {
	s.subjects = append(s.subjects, subs...)
}

func (s *students) AddGrades(grades map[string][]int) {
	for name, gs := range grades {
		for _, v := range gs {
			s.s[name].grades = append(s.s[name].grades, v)
			s.s[name].totalMarks += v
		}
	}
}

func (s *students) PrintDefault(w io.Writer) {
	table := newTable(w)
	table.SetHeader(s.headers())

	names := s.sortedStudentNames(SortOrder(2), len(s.subjects))
	for i, name := range names {
		row := []string{strconv.Itoa(i + 1), name}
		student := s.s[name]
		if len(student.grades) > 0 {
			row = append(row, student.gradesPrint()...)
			row = append(row, strconv.Itoa(student.totalMarks))
		}
		table.Append(row)
	}

	table.Render()
}

func (s *students) Print(w io.Writer, o SortOrder, subj int, st SortType) {
	table := newTable(w)
	table.SetHeader(s.headers())

	names := s.sortedNamesBySubj(o, subj, st)
	for i, name := range names {
		row := []string{strconv.Itoa(i + 1), name}
		student := s.s[name]
		if len(student.grades) > 0 {
			row = append(row, student.gradesPrint()...)
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

func (s *students) sortedNamesBySubj(o SortOrder, subj int, st SortType) []string {
	var names []string
	for k := range s.s {
		names = append(names, k)
	}

	switch st {
	case QuickSort:
		names = s.quicksort(names, o, subj)
	case MergeSort:
		names = s.mergesort(names, o, subj)
	case SelectionSort:
		names = s.selectionsort(names, o, subj)
	}

	return names
}

func (s *students) sortedStudentNames(o SortOrder, subj int) []string {
	var names []string
	for k := range s.s {
		names = append(names, k)
	}

	return s.quicksort(names, o, subj)
}

func (s *students) quicksort(n []string, o SortOrder, subj int) []string {
	if len(n) < 2 {
		return n
	}

	p := s.getPivot(n, subj)
	pivStudent := s.s[n[p]]
	n = swap(n, p)

	var lastMoved int
	for i := 0; i < len(n)-1; i++ {
		curStudent := s.s[n[i]]
		if compare(o, curStudent.getMarkByType(subj), pivStudent.getMarkByType(subj)) {
			n[lastMoved], n[i] = n[i], n[lastMoved]
			lastMoved++
		}
	}
	n = swap(n, lastMoved)
	out := append(s.quicksort(n[:lastMoved], o, subj), n[lastMoved])
	out = append(out, s.quicksort(n[lastMoved+1:], o, subj)...)
	return out
}

func (s *students) getPivot(n []string, subj int) int {
	first := s.s[n[0]].getMarkByType(subj)
	mid := s.s[n[len(n)/2]].getMarkByType(subj)
	last := s.s[n[len(n)-1]].getMarkByType(subj)

	if first > mid && first < last {
		return 0
	} else if mid > last {
		return len(n) - 1
	}

	return len(n) / 2
}

func swap(n []string, i int) []string {
	n[len(n)-1], n[i] = n[i], n[len(n)-1]
	return n
}

func (s *students) mergesort(n []string, o SortOrder, subj int) []string {
	if len(n) < 2 {
		return n
	}
	mid := len(n) / 2
	return s.merge(s.mergesort(n[:mid], o, subj), s.mergesort(n[mid:], o, subj), o, subj)
}

func (s *students) merge(lArr, rArr []string, o SortOrder, subj int) []string {
	sz := len(lArr) + len(rArr)
	n := make([]string, sz, sz)

	var i, j int
	for k := 0; k < sz; k++ {
		if i > len(lArr)-1 && j < len(rArr) {
			n[k] = rArr[j]
			j++
			continue
		} else if j > len(rArr)-1 && i < len(lArr) {
			n[k] = lArr[i]
			i++
			continue
		}

		lStud, rStud := s.s[lArr[i]], s.s[rArr[j]]
		if compare(o, lStud.getMarkByType(subj), rStud.getMarkByType(subj)) {
			n[k] = lArr[i]
			i++
		} else {
			n[k] = rArr[j]
			j++
		}
	}

	return n
}

func (s *students) selectionsort(n []string, o SortOrder, subj int) []string {
	for i := 0; i < len(n)-1; i++ {
		curMark := s.s[n[i]].getMarkByType(subj)
		low := curMark
		lowIdx := i
		for j := i + 1; j < len(n); j++ {
			nMark := s.s[n[j]].getMarkByType(subj)
			if compare(o, nMark, curMark) && compare(o, nMark, low) {
				low = nMark
				lowIdx = j
			}
		}
		if lowIdx != i {
			n[i], n[lowIdx] = n[lowIdx], n[i]
		}
	}
	return n
}

type student struct {
	name       string
	grades     []int
	totalMarks int
}

func (s *student) getMarkByType(subject int) int {
	if subject == len(s.grades) {
		return s.totalMarks
	}
	return s.grades[subject]
}

func (s *student) gradesPrint() []string {
	o := make([]string, 0, len(s.grades))
	for _, v := range s.grades {
		o = append(o, strconv.Itoa(v))
	}
	return o
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

func compare(o SortOrder, a, b int) bool {
	if o == ASC {
		return a < b
	}
	return a > b
}
