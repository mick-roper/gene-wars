package core

import "testing"

func Test_NewPopulation(t *testing.T) {
	size := 10

	p := NewPopulation(size)

	if p == nil {
		t.Error("population is nil")
		return
	}

	if p.Fittest != 0 {
		t.Error("expected 'Fittest' to be 0")
		return
	}

	if len(p.Individuals) != size {
		t.Errorf("expected array of individuals to havel length of %v", size)
		return
	}

	for i := 0; i < len(p.Individuals); i++ {
		if p.Individuals[i] == nil {
			t.Errorf("inidividual at index %v is nil", i)
		}
	}
}
