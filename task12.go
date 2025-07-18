package main

type Set[K comparable] struct {
	m map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{
		m: make(map[K]struct{}),
	}
}

func (s *Set[K]) Add(e K) {
	s.m[e] = struct{}{}
}

func (s *Set[K]) GetAll() []K {
	result := make([]K, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}

func MakeUnique(numbers []string) []string {
	set := NewSet[string]()
	for _, n := range numbers {
		set.Add(n)
	}
	return set.GetAll()
}
