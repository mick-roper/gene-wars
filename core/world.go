package core

// World that a population lives in
type World struct {
	Population *Population
	Generation int
}

// NewWorld creates a new World
func NewWorld(popSize, genes int) *World {
	p := NewPopulation(popSize, genes)

	return &World{p, 0}
}
