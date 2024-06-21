package game_test

import (
	"fmt"
	"testing"

	"github.com/jcpage573/prisoners-dilemma/internal/game"
	"github.com/jcpage573/prisoners-dilemma/internal/prisoner"
	"gonum.org/v1/gonum/stat/combin"
)

func createTestPrisoners() []prisoner.Prisoner {
	var pl []prisoner.Prisoner
	for i := 0; i < 4000; i++ {
		pl = append(pl, prisoner.Prisoner{Name: fmt.Sprint("Joe ", i)})
	}
	return pl
}

func TestWomboCombo(t *testing.T) {
	var playerList = createTestPrisoners()
	var choose2 = combin.Binomial(len(playerList), 2)
	if len(game.WomboCombo(playerList)) != choose2 {
		t.Error("Combos aint right")
	}
}

func TestGonumCombo(t *testing.T) {
	combin.Permutations(4000, 2)
}
