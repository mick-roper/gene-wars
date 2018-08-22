package core

import (
	"math/rand"
	"time"
)

// Individual that contains genes
type Individual struct {
	Fitness int
	genes   []int
}

// NewIndividual creates a new individual
func NewIndividual(geneLength int) *Individual {
	genes := make([]int, geneLength)

	src := rand.NewSource(time.Now().UnixNano())
	x := rand.New(src).Int()

	for i := 0; i < len(genes); i++ {
		genes[i] = x & 0x01
		x >>= 1
	}

	return &Individual{0, genes}
}

// CalcFitness calculates the fitness of the individual
func (i *Individual) CalcFitness() {
	f := 0

	for n := 0; n < len(i.genes); n++ {
		if i.genes[n] == 1 {
			f++
		}
	}

	i.Fitness = f
}
