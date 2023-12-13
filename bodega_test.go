package main

import (
	"math/rand"
	"testing"
)

func PrepareRealSet(limit int) map[uint32]bool {
	realSet := make(map[uint32]bool)

	for i := 0; i < limit; i++ {
		sample := rand.Uint32()
		realSet[sample] = true
	}

	return realSet
}

func TestAccurateLearningModel(t *testing.T) {
	realSet := PrepareRealSet(100)
	lm := NewLearningModel(realSet, 1.0)

	for value, _ := range realSet {
		if !lm.Test(value) {
			t.Fatalf(`All test values should return true: value %v`, value)
		}
	}
}

func TestInaccurateLearningModel(t *testing.T) {
	limit := 1000

	realSet := PrepareRealSet(limit)
	secondSet := PrepareRealSet(limit)

	lm := NewLearningModel(realSet, 0.0)

	incorrect := 0

	for value, _ := range realSet {
		if !lm.Test(value) {
			incorrect += 1
		}
	}

	t.Log("False negatives:", incorrect, " Expected: ~", limit/2)

	incorrect = 0

	for value, _ := range secondSet {
		if lm.Test(value) {
			incorrect += 1
		}
	}

	t.Log("False positives:", incorrect, " Expected: ~", limit/2)
}
