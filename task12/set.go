package main

// Set - множество, состоящее из уникальных элементов типа V.
// Не поддерживает конкурентность.
type Set[V comparable] struct {
	// Будем хранить элементы в качестве ключей карты, чтобы обеспечить поиск элементов
	// в множестве за O(n).
	m map[V]bool
}

func NewSet[V comparable](elements ...V) *Set[V] {
	s := &Set[V]{
		m: make(map[V]bool, len(elements)),
	}

	for _, e := range elements {
		s.Add(e)
	}

	return s
}

// Has проверяет, есть ли в множестве элемент v.
func (s *Set[V]) Has(v V) bool {
	_, ok := s.m[v]
	return ok
}

// Add добавляет в множество элемент v.
func (s *Set[V]) Add(v V) {
	s.m[v] = true
}

// Walk позволяет пройти по всем элементам множества.
func (s *Set[V]) Walk(walker func(v V)) {
	for v := range s.m {
		walker(v)
	}
}
