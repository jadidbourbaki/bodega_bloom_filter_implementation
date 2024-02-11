package main

import "testing"

func TestPseudorandomPermutationCorrectness(t *testing.T) {
	prp := NewPseudorandomPermutation(PSEUDORANDOM_TEST_KEY, PSEUDORANDOM_TEST_IV)

	var value uint32 = 42

	if prp.Decrypt(prp.Encrypt(value)) != value {
		t.Fatal("Incorrect PRP Encryption / Decryption")
	} else {
		t.Log("PRP Encryption / Decryption test passes")
	}
}
