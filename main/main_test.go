package main

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

var playerList = createTestPrisoners()
var choose2 = combin.Binomial(len(playerList), 2)

func createTestPrisoners() []Prisoner {
	var pl []Prisoner
	for i := 0; i < 4000; i++ {
		pl = append(pl, Prisoner{name: fmt.Sprint("Joe ", i)})

	}
	return pl
}

func TestWomboCombo(t *testing.T) {
	if len(WomboCombo(playerList)) != choose2 {
		t.Error("Combos aint right")
	}
}

func TestGonumCombo(t *testing.T) {
	combin.Permutations(len(playerList), 2)
}
