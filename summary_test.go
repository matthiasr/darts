package main

import (
	"reflect"
	"testing"
)

func (s summary) assertState(t *testing.T, count, sum uint, hist map[uint]uint) {
	if s.count != count {
		t.Errorf("expected %+v to have count %d", s, count)
	}
	if s.sum != sum {
		t.Errorf("expected %+v to have sum %d", s, sum)
	}
	// If no histogram provided, don't check
	if hist == nil {
		return
	}
	if !reflect.DeepEqual(s.hist, hist) {
		t.Errorf("expected histogram %v, got %v", hist, s.hist)
	}
}

func TestObserve(t *testing.T) {
	s := summary{}
	s.Observe(2)
	s.assertState(t, 1, 2, map[uint]uint{2: 1})
	s.Observe(1)
	s.assertState(t, 2, 3, map[uint]uint{2: 1, 1: 1})
	s.Observe(2)
	s.assertState(t, 3, 5, map[uint]uint{2: 2, 1: 1})
}

func (s summary) assertAverage(t *testing.T, want float64) {
	got := s.Average()
	if got < want-0.001 || got > want+0.001 {
		t.Errorf("expected %+v to have average %v, got %v", s, want, got)
	}
}

func TestAverage(t *testing.T) {
	s := summary{}
	s.Observe(2)
	s.assertAverage(t, 2)
	s.Observe(1)
	s.assertAverage(t, 1.5)
	s.Observe(2)
	s.assertAverage(t, 1.66667)
}
