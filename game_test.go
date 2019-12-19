package main

import "testing"

import "math"

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

// fixedHitter returns a hitter that always hits the same point
func fixedHitter(x, y float64) hitter {
	return func() (float64, float64) {
		return x, y
	}
}

// compareGames is true if the games are logically equivalent: the hitter never
// matters; the radius is only relevant unless the game is already finished.
func compareGames(this, that Game) bool {
	if this.finished {
		return that.finished && this.score == that.score
	}

	switch {
	case this.finished != that.finished:
		return false
	case this.score != that.score:
		return false
	case 0.99 > (this.radius / that.radius):
		return false
	case 1.01 < (this.radius / that.radius):
		return false
	default:
		return true
	}
}

type stepCase struct {
	current, next Game
}

func TestStep(t *testing.T) {
	cases := []stepCase{
		stepCase{
			Game{radius: 1.0, hitter: fixedHitter(1.0, 1.0)},
			Game{finished: true, score: 1},
		},
		stepCase{
			Game{radius: 1.0, hitter: fixedHitter(0.5, 0.5)},
			Game{radius: math.Sqrt(0.5), finished: false, score: 1},
		},
		stepCase{
			Game{radius: math.Sqrt(0.5), hitter: fixedHitter(0.6, 0.5), score: 1},
			Game{finished: true, score: 2},
		},
	}

	for _, c := range cases {
		if got := c.current.Step(); !compareGames(c.next, got) {
			t.Errorf("expected step:\n  %s -> %s\ngot:\n  %s -> %s", c.current, c.next, c.current, got)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGame().Run()
	}
}
