package main

import (
	"math/rand"
	"testing"
)

func TestDowntownBodegaFilterCorrectness(t *testing.T) {
	realSet := make(map[uint32]bool)

	for i := 0; i < 5; i++ {
		sample := rand.Uint32()
		realSet[sample] = true
	}

	downtown := NewDowntownBodegaFilter(90, 3, 10, 3, realSet, 0.9, PSEUDORANDOM_TEST_KEY, PSEUDORANDOM_TEST_IV, PSEUDORANDOM_TEST_KEY, PSEUDORANDOM_TEST_IV)

	for value, _ := range realSet {
		if !downtown.Test(value) {
			t.Fatal("Correctness property of Downtown Bodega Filter fails")
		}
	}

}
