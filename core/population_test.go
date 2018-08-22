package core

import (
	"reflect"
	"testing"
)

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

func Test_GetFittest(t *testing.T) {
	i := []*Individual{
		&Individual{0, []int{1, 1, 1, 0, 0}},
		&Individual{0, []int{1, 0, 1, 0, 0}},
		&Individual{0, []int{1, 1, 1, 1, 0}}, // should be fittest!
		&Individual{0, []int{1, 0, 0, 0, 0}},
	}

	for _, x := range i {
		x.CalcFitness()
	}

	p := &Population{i, 0}

	expect := i[2]
	fittest := p.GetFittest()

	if p.Fittest != 2 {
		t.Error("expected fittest to be 2")
	}

	if !reflect.DeepEqual(fittest, expect) {
		t.Errorf("expected %v to be %v", fittest, expect)
	}
}
