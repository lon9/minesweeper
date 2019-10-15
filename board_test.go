package main

import "testing"

func TestInitializeBoard(t *testing.T) {
	b := NewBoard(Easy)
	b.showBoard()
}
