package core

import (
	"math"
)

// Population of individuals
type Population struct {
	Individuals []*Individual
	Fittest     int
}

// NewPopulation creates a new population of a given size
func NewPopulation(popSize, genes int) *Population {
	individuals := make([]*Individual, popSize)

	for i := 0; i < len(individuals); i++ {
		individuals[i] = NewIndividual(genes)
	}

	return &Population{individuals, 0}
}

// GetFittest returns the fittest individual
func (p *Population) GetFittest() *Individual {
	maxFit := math.MinInt32
	maxFitIx := 0

	for i := 0; i < len(p.Individuals); i++ {
		x := p.Individuals[i]
		if maxFit <= x.Fitness {
			maxFit = x.Fitness
			maxFitIx = i
		}
	}

	p.Fittest = maxFitIx
	return p.Individuals[maxFitIx]
}

// GetSecondFittest returns the second fittest individual
func (p *Population) GetSecondFittest() *Individual {
	maxFit1 := 0
	maxFit2 := 0

	for i := 0; i < len(p.Individuals); i++ {
		if p.Individuals[i].Fitness > p.Individuals[maxFit1].Fitness {
			maxFit2 = maxFit1
			maxFit1 = i
		} else if p.Individuals[i].Fitness > p.Individuals[maxFit2].Fitness {
			maxFit2 = i
		}
	}

	return p.Individuals[maxFit2]
}

// GetLeastFittestIndex returns the index of the least fittest individual in the population
func (p *Population) GetLeastFittestIndex() int {
	minFit := math.MaxInt32
	minFitIx := 0

	for i := 0; i < len(p.Individuals); i++ {
		if minFit >= p.Individuals[i].Fitness {
			minFit = p.Individuals[i].Fitness
			minFitIx = i
		}
	}

	return minFitIx
}

// CalculatePopulationFitness calculate the fitness of the entire population
func (p *Population) CalculatePopulationFitness() {
	for i := 0; i < len(p.Individuals); i++ {
		p.Individuals[i].CalcFitness()
	}

	p.GetFittest()
}
