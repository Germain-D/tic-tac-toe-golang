package main

import (
	"testing"
)

func TestWinner(t *testing.T) {
	terrain := map[int]string{
		1: "X", 2: " ", 3: " ",
		4: " ", 5: "X", 6: " ",
		7: " ", 8: " ", 9: "X",
	}
	terrain1 := map[int]string{
		1: "O", 2: " ", 3: " ",
		4: "O", 5: "X", 6: " ",
		7: "O", 8: " ", 9: "X",
	}
	terrain3 := map[int]string{
		1: "O", 2: " ", 3: " ",
		4: "X", 5: "X", 6: "X",
		7: "O", 8: " ", 9: "X",
	}
	tour := 1
	if winner(terrain, tour) != 10 {
		t.Error("Expected 10, got ", winner(terrain, tour))
	}
	if winner(terrain1, tour) != 10 {
		t.Error("Expected 10, got ", winner(terrain1, tour))
	}
	if winner(terrain3, tour) != 10 {
		t.Error("Expected 10, got ", winner(terrain3, tour))
	}
}
