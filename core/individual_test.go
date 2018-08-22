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
