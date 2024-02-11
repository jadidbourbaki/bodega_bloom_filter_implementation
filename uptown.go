// main package
package main

import (
	"github.com/bits-and-blooms/bloom/v3"
)

type UptownBodegaFilter struct {
	prp                PseudorandomPermutation
	bloomAbove         bloom.BloomFilter
	learningModelPatty LearningModel
	bloomBelow         bloom.BloomFilter
}

func NewUptownBodegaFilter(bitsAbove uint, hashesAbove uint, bitsBelow uint, hashesBelow uint,
	realSet map[uint32]bool, learningModelAccuracy float64, key []byte, iv []byte) *UptownBodegaFilter {
	bloomAbove := bloom.New(bitsAbove, hashesAbove)
	bloomBelow := bloom.New(bitsBelow, hashesBelow)
	learningModelPatty := NewLearningModel(realSet, learningModelAccuracy)
	prp := NewPseudorandomPermutation(key, iv)

	for value, _ := range realSet {
		serialized := prp.Encrypt(value)
		bloomAbove.Add(serialized)
		bloomBelow.Add(serialized)
	}

	uptown := UptownBodegaFilter{bloomAbove: *bloomAbove, bloomBelow: *bloomBelow, learningModelPatty: *learningModelPatty, prp: *prp}
	return &uptown
}

func (uptown *UptownBodegaFilter) Test(value uint32) bool {
	serialized := uptown.prp.Encrypt(value)

	if !uptown.bloomAbove.Test(serialized) {
		return false
	}

	if uptown.learningModelPatty.Test(value) {
		return true
	}

	return uptown.bloomBelow.Test(serialized)
}
