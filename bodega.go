// main package
package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/bits-and-blooms/bloom/v3"
)

func serialize(value uint32) []byte {
	rtn := make([]byte, 4)
	binary.BigEndian.PutUint32(rtn, value)
	return rtn
}

type LearningModel struct {
	realSet          map[uint32]bool
	correctSet       map[uint32]bool
	falsePositiveSet map[uint32]bool
	falseNegativeSet map[uint32]bool
	accuracy         float64
}

func NewLearningModel(realSet map[uint32]bool, accuracy float64) *LearningModel {
	lm := LearningModel{realSet: realSet, accuracy: accuracy}

	lm.correctSet = make(map[uint32]bool)
	lm.falsePositiveSet = make(map[uint32]bool)
	lm.falseNegativeSet = make(map[uint32]bool)

	return &lm
}

func (lm *LearningModel) Test(value uint32) bool {
	_, realOk := lm.realSet[value]

	_, correctOk := lm.correctSet[value]

	if correctOk {
		return realOk
	}

	_, falseNegativeOk := lm.falseNegativeSet[value]

	if falseNegativeOk {
		return false
	}

	_, falsePositiveOk := lm.falsePositiveSet[value]

	if falsePositiveOk {
		return true
	}

	if rand.Float64() <= lm.accuracy {
		lm.correctSet[value] = true
		return realOk
	}

	falseNegative := ((rand.Uint32() % 2) == 0)

	if falseNegative {
		lm.falseNegativeSet[value] = true
		return false
	}

	// Otherwise, false positive
	lm.falsePositiveSet[value] = true
	return true
}

type BodegaBloomFilter struct {
	bloomAbove         bloom.BloomFilter
	learningModelPatty LearningModel
	bloomBelow         bloom.BloomFilter
}

func NewBodegaBloomFilter(bitsAbove uint, hashesAbove uint, bitsBelow uint, hashesBelow uint,
	realSet map[uint32]bool, learningModelAccuracy float64) *BodegaBloomFilter {
	bloomAbove := bloom.New(bitsAbove, hashesAbove)
	bloomBelow := bloom.New(bitsBelow, hashesBelow)
	learningModelPatty := NewLearningModel(realSet, learningModelAccuracy)

	for value, _ := range realSet {
		serialized := serialize(value)
		bloomAbove.Add(serialized)
		bloomBelow.Add(serialized)
	}

	bodega := BodegaBloomFilter{bloomAbove: *bloomAbove, bloomBelow: *bloomBelow, learningModelPatty: *learningModelPatty}
	return &bodega
}

func (bodega *BodegaBloomFilter) Test(value uint32) bool {
	serialized := serialize(value)

	if !bodega.bloomAbove.Test(serialized) {
		return false
	}

	if bodega.learningModelPatty.Test(value) {
		return true
	}

	return bodega.bloomBelow.Test(serialized)
}

func main() {
	filter := bloom.New(100, 3)

	realSet := make(map[uint32]bool)

	for i := 0; i < 5; i++ {
		sample := rand.Uint32()
		realSet[sample] = true
	}

	for value, _ := range realSet {
		serialized := serialize(value)
		filter.Add(serialized)
	}

	lm := NewLearningModel(realSet, 0.5)
	bodega := NewBodegaBloomFilter(10, 3, 10, 3, realSet, 0.9)

	for value, _ := range realSet {
		serialized := serialize(value)
		fmt.Println("Bloom Filter: ", filter.Test(serialized))
		fmt.Println("Learning Model: ", lm.Test(value))
		fmt.Println("Bodega Bloom Filter: ", bodega.Test(value))
	}
}
