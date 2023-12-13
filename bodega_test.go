package main

import (
	"math/rand"
	"testing"
)

func TestAccurateLearningModel(t *testing.T) {
	realSet := make(map[uint32]bool)

	for i := 0; i < 100; i++ {
		sample := rand.Uint32()
		realSet[sample] = true
	}

	lm := NewLearningModel(realSet, 1.0)

	for value, _ := range realSet {
		if !lm.Test(value) {
			t.Fatalf(`All test values should return true: value %v`, value)
		}
	}
}
