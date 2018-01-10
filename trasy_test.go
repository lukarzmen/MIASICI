package main

import "testing"

func TestTrasy(t *testing.T) {
	if len(trasy) < 1 {
		t.Errorf("Brak tras w rutingu")
	}
}
