package main

import "math/rand"

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
