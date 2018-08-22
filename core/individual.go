package core

const geneCount = 5

// Individual that contains genes
type Individual struct {
	fitness float64
	genes   []float64
}

// NewIndividual creates a new individual
func NewIndividual() *Individual {
	return &Individual{
		fitness: 0,
		genes:   make([]float64, geneCount),
	}
}
