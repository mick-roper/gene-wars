package core

import "testing"

func Test_NewWorld(t *testing.T) {
	const popSize = 10
	const genes = 5

	w := NewWorld(popSize, genes)

	if w == nil {
		t.Error("world is nil")
		return
	}
}
