package solutions

import "fmt"

type Set struct {
	m   map[string]bool
	set []string
}

func NewSetFromSlice(l []string) *Set {
	// в множестве проверяем уникальность с помощью map
	set := make([]string, 0, 100)
	m := make(map[string]bool, 100)
	for _, v := range l {
		_, ok := m[v]
		if !ok {
			m[v] = true
			set = append(set, v)
		}
	}
	return &Set{
		set: set,
		m:   m,
	}
}

func NewSet() *Set {
	return &Set{
		set: make([]string, 0, 100),
		m:   make(map[string]bool, 100),
	}
}

func (s *Set) String() string {
	return fmt.Sprintf("%v", s.set)
}

func (s *Set) Append(v string) {
	_, ok := s.m[v]
	if !ok {
		s.set = append(s.set, v)
	}
}

func (s *Set) Remove(v string) {
	_, ok := s.m[v]
	if ok {
		for i, str := range s.set {
			if str == v {
				s.set = append(s.set[:i], s.set[i+1:]...)
				break
			}
		}
	}
}

func Example12() {
	l := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSetFromSlice(l)
	fmt.Println(set)
	set.Append("table")
	set.Append("cat")
	fmt.Println(set)
	set.Remove("tree")
	set.Remove("ag")
	fmt.Println(set)
}
