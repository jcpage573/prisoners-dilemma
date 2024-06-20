package src

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

func createTestPrisoners() []Prisoner {
	var pl []Prisoner
	for i := 0; i < 4000; i++ {
		pl = append(pl, Prisoner{Name: fmt.Sprint("Joe ", i)})
	}
	return pl
}

func TestWomboCombo(t *testing.T) {
	var playerList = createTestPrisoners()
	var choose2 = combin.Binomial(len(playerList), 2)
	if len(WomboCombo(playerList)) != choose2 {
		t.Error("Combos aint right")
	}
}

func TestGonumCombo(t *testing.T) {
	combin.Permutations(4000, 2)
}
