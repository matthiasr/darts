package main

type summary struct {
	count uint
	sum   uint
	hist  map[uint]uint
}

func (s *summary) Observe(value uint) {
	s.count++
	s.sum += value
	if s.hist == nil {
		s.hist = make(map[uint]uint)
	}
	s.hist[value]++
}

func (s *summary) Average() float64 {
	return float64(s.sum) / float64(s.count)
}
