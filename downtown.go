package main

import "github.com/bits-and-blooms/bloom/v3"

/**


       -------    Negatives
   --> | LM  | --------------- PRP B -- Bloom B
       -------
          | Positives
          |
	 PRP A
	  |
	 Bloom A
**/

type DowntownBodegaFilter struct {
	prpA               PseudorandomPermutation
	prpB               PseudorandomPermutation
	bloomA             bloom.BloomFilter
	bloomB             bloom.BloomFilter
	learningModelAbove LearningModel
}

func NewDowntownBodegaFilter(bitsA uint, hashesA uint, bitsB uint, hashesB uint,
	realSet map[uint32]bool, learningModelAccuracy float64, keyA []byte, ivA []byte,
	keyB []byte, ivB []byte) *DowntownBodegaFilter {

	bloomA := bloom.New(bitsA, hashesA)
	bloomB := bloom.New(bitsB, hashesB)

	learningModelAbove := NewLearningModel(realSet, learningModelAccuracy)

	prpA := NewPseudorandomPermutation(keyA, ivA)
	prpB := NewPseudorandomPermutation(keyB, ivB)

	// Might need to check this logic
	for value, _ := range realSet {
		serializedA := prpA.Encrypt(value)
		serializedB := prpB.Encrypt(value)

		bloomA.Add(serializedA)
		bloomB.Add(serializedB)
	}

	downtown := DowntownBodegaFilter{bloomA: *bloomA, bloomB: *bloomB, learningModelAbove: *learningModelAbove, prpA: *prpA, prpB: *prpB}
	return &downtown
}

func (downtown *DowntownBodegaFilter) Test(value uint32) bool {
	serializedA := downtown.prpA.Encrypt(value)
	serializedB := downtown.prpB.Encrypt(value)

	if downtown.learningModelAbove.Test(value) {
		return downtown.bloomA.Test(serializedA)
	}

	return downtown.bloomB.Test(serializedB)
}
