// main package
package main

import (
	"encoding/binary"

	"github.com/bits-and-blooms/bloom/v3"
)

func serialize(value uint32) []byte {
	rtn := make([]byte, 4)
	binary.BigEndian.PutUint32(rtn, value)
	return rtn
}

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

	bodega := UptownBodegaFilter{bloomAbove: *bloomAbove, bloomBelow: *bloomBelow, learningModelPatty: *learningModelPatty, prp: *prp}
	return &bodega
}

func (bodega *UptownBodegaFilter) Test(value uint32) bool {
	serialized := bodega.prp.Encrypt(value)

	if !bodega.bloomAbove.Test(serialized) {
		return false
	}

	if bodega.learningModelPatty.Test(value) {
		return true
	}

	return bodega.bloomBelow.Test(serialized)
}
