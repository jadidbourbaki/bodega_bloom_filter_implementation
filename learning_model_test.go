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

func TestLearningModelAccurate(t *testing.T) {
	realSet := PrepareRealSet(100)
	lm := NewLearningModel(realSet, 1.0)

	for value, _ := range realSet {
		if !lm.Test(value) {
			t.Fatalf(`All test values should return true: value %v`, value)
		}
	}
}

func TestLearningModelInaccurate(t *testing.T) {
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

func TestLearningModelPersistence(t *testing.T) {
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

	check_again := 0

	for value, _ := range realSet {
		if !lm.Test(value) {
			check_again += 1
		}
	}

	if incorrect != check_again {
		t.Fatal("False negatives not persistent")
	} else {
		t.Log("False negatives are persistent")
	}

	incorrect = 0

	for value, _ := range secondSet {
		if lm.Test(value) {
			incorrect += 1
		}
	}

	check_again = 0

	for value, _ := range secondSet {
		if lm.Test(value) {
			check_again += 1
		}
	}

	if incorrect != check_again {
		t.Fatal("False positives not persistent")
	} else {
		t.Log("False positives are persistent")
	}
}
