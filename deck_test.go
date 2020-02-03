package main

import "testing"

// t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but received %v", len(d))
	}
}
