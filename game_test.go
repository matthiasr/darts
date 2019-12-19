package main

import "testing"

type isHitTestCase struct {
	r, x, y float64
	hit     bool
}

func TestIsHit(t *testing.T) {
	cases := []isHitTestCase{
		isHitTestCase{1.0, 0.99, 0.99, false},
		isHitTestCase{1.0, 0.0, 0.0, true},
		isHitTestCase{0.5, 0.4, 0.0, true},
		isHitTestCase{0.5, -0.1, 0.2, true},
		isHitTestCase{0.1, 0.1, -0.1, false},
	}

	for i, c := range cases {
		g := Game{radius: c.r}
		if want, got := c.hit, g.isHit(c.x, c.y); want != got {
			t.Errorf("case %d: expected %v.isHit(%v,%v) to be %v, got %v", i, g, c.x, c.y, want, got)
		}
	}
}
