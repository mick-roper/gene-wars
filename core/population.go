package core

// Population of individuals
type Population struct {
	Individuals []*Individual
	Fittest     int
}

// NewPopulation creates a new population of a given size
func NewPopulation(popSize int) *Population {
	const genes = 10
	individuals := make([]*Individual, popSize)

	for i := 0; i < len(individuals); i++ {
		individuals[i] = NewIndividual(genes)
	}

	return &Population{individuals, 0}
}
