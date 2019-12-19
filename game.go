package main

import "math/rand"

import "fmt"

type hitter func() (float64, float64)

// TODO(mr): is it sound to only work in the top-right quarter?
func randomHitter() (float64, float64) {
	return rand.Float64(), rand.Float64()
}

// Game represents the state of a game, including a hit generator.
type Game struct {
	hitter   hitter
	radius   float64
	score    uint
	finished bool
}

// NewGame initializes a new game using the standard random hitter.
func NewGame() *Game {
	return &Game{
		hitter:   randomHitter,
		radius:   1.0,
		score:    0, // score is incremented one last time in the last throw
		finished: false,
	}
}

func (g *Game) isHit(x, y float64) bool {
	return x*x+y*y < g.radius*g.radius
}

// Step iterates the game once.
func (g *Game) Step() {
	if g.finished {
		return
	}

	g.score++
	x, y := g.hitter()
	if g.isHit(x, y) {
		return
	}
	g.finished = true
	return
}

// Run a single game to completion, returning the final score.
func (g *Game) Run() uint {
	for !g.finished {
		g.Step()
	}
	return g.score
}

// String returns a readable representation of the game.
func (g *Game) String() string {
	return fmt.Sprintf("finished=%v,radius=%v,score=%d", g.finished, g.radius, g.score)
}
