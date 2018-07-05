package pkg

type Set struct {
	values map[string]struct{}
}

func NewSet() *Set {
	return &Set{values: make(map[string]struct{})}
}

func (s Set) Add(v string)  {
	s.values[v] = struct{}{}
}

func (s Set) AddAll(slice []string)  {
	for _, v := range slice {
		s.Add(v)
	}
}

func (s Set) Remove(v string) {
	delete(s.values, v)
}

func (s Set) Contains(v string) bool {
	_, found := s.values[v]
	return found
}

func (s Set) Size() int {
	return len(s.values)
}