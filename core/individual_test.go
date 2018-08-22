package core

import (
	"testing"
)

func Test_NewIndividual(t *testing.T) {
	const l = 10
	i := NewIndividual(l)

	if i == nil {
		t.Error("individual is nil")
		return
	}

	if i.Fitness != 0 {
		t.Error("fitness should be 0")
		return
	}

	if len(i.genes) != l {
		t.Error("genes array length is wrong")
	}
}

func Test_CalcFitness(t *testing.T) {
	i := NewIndividual(10)

	expectFitness := 0

	for n := 0; n < len(i.genes); n++ {
		if i.genes[n] == 1 {
			expectFitness++
		}
	}

	i.CalcFitness()

	if expectFitness != i.Fitness {
		t.Errorf("expected fitness to be %v but was %v", expectFitness, i.Fitness)
	}
}
