package core

import (
	"reflect"
	"testing"
)

func Test_NewIndividual(t *testing.T) {
	tests := []struct {
		name   string
		expect *Individual
	}{
		{
			name:   "all good",
			expect: &Individual{fitness: 0, genes: make([]float64, GeneCount)},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i := NewIndividual()

			if !reflect.DeepEqual(i, test.expect) {
				t.Errorf("expected %v to be %v", i, test.expect)
			}
		})
	}
}
