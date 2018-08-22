package core

import (
	"testing"
)

func Test_NewIndividual(t *testing.T) {
	i := NewIndividual()

	if i == nil {
		t.Error("individual is nil")
		return
	}

	if len(i.genes) != GeneCount {
		t.Error("genes array length is wrong")
	}
}
