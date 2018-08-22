package core

import (
	"math/rand"
	"time"
)

// GeneCount indicates the number of genes in each individual
const GeneCount = 5

// Individual that contains genes
type Individual struct {
	fitness int
	genes   []int
}

// NewIndividual creates a new individual
func NewIndividual() *Individual {
	genes := make([]int, GeneCount)

	src := rand.NewSource(time.Now().UnixNano())
	x := rand.New(src).Int()

	for i := 0; i < GeneCount; i++ {
		genes[i] = x & 0x01
		x >>= 1
	}

	return &Individual{0, genes}
}
