package main

// SetUnion define a set
type SetUnion struct {
	myset    []int
	myweight []int
}

func (s *SetUnion) create(size int) {
	s.myset = make([]int, size)
	s.myweight = make([]int, size)

	for i := 0; i < len(s.myset); i++ {
		s.myset[i] = i
	}
}

func (s *SetUnion) union(a, b int) {
	roota := s.root(a)
	rootb := s.root(b)

	if roota == rootb {
		return
	}

	if s.myweight[roota] < s.myweight[rootb] {
		s.myset[roota] = rootb
		s.myweight[rootb] += s.myweight[roota]
		return
	}

	s.myset[rootb] = roota
	s.myweight[roota] += s.myweight[rootb]
}

func (s *SetUnion) isConnected(a, b int) bool {
	return s.root(a) == s.root(b)
}

func (s *SetUnion) root(i int) int {
	for i != s.myset[i] {
		// flattening the tree
		// change my actual root to be the root of my root
		myroot := s.myset[i]
		s.myset[i] = s.myset[myroot]
		i = s.myset[i]
	}

	return i
}
