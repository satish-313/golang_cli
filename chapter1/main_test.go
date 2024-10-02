package main

import (
	"bytes"
	"testing"
)

func TestCountWord(t *testing.T) {
	b := bytes.NewBufferString("word1 \n word2 word3 word4\n")

	exp := 4
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("word1 \n word2 \n word3 \n ")

	exp := 4
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expecting %d, got %d instead.\n", exp, res)
	}
}

func TestByteCount(t *testing.T) {
	b := bytes.NewBufferString("total number of bytes")

	exp := b.Len()
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expecting %d, got %d instead.\n", exp, res)
	}
}
