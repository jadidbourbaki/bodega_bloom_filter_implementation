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
	realSet  map[uint32]bool
	accuracy float64
}

func NewLearningModel(realSet map[uint32]bool, accuracy float64) *LearningModel {
	lm := LearningModel{realSet: realSet, accuracy: accuracy}
	return &lm
}

func (lm *LearningModel) Test(value uint32) bool {
	if rand.Float64() > lm.accuracy {
		return ((rand.Uint32() % 2) == 0)
	}

	_, ok := lm.realSet[value]

	return ok
}

type BodegaBloomFilter struct {
	firstBloom  bloom.BloomFilter
	secondBloom bloom.BloomFilter
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

	lm := NewLearningModel(realSet, 0.9)

	for value, _ := range realSet {
		serialized := serialize(value)
		fmt.Println("Bloom Filter: ", filter.Test(serialized))
		fmt.Println("Learning Model: ", lm.Test(value))

	}
}
