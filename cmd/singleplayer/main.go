package main

import (
	"fmt"
	"math/rand"

	pd "github.com/jcpage573/prisoners-dilemma/src"
)

// Example Strategies
// oa: own answers
// opa: opponent answers
// return `true` to cooperate, `false` to defect

func random(oa []bool, opa []bool) bool {
	return rand.Intn(2) == 1
}

func defectCooperateTwiceThenDefect(oa []bool, opa []bool) bool {
	return len(oa)%3 != 0
}

func titForTat(oa []bool, opa []bool) bool {
	n := len(oa)
	if n < 1 {
		return true
	}

	return opa[n-1]
}

func CreatePrisoners() []pd.Prisoner {
	prisoners := []pd.Prisoner{
		{Name: "Random", Owner: "Jackson", Strategy: random},
		{Name: "Mod 3 Prisoner", Owner: "Jackson", Strategy: defectCooperateTwiceThenDefect},
		{Name: "Tit for Tat", Owner: "Hunter", Strategy: titForTat},
	}
	return prisoners
}

func main() {
	prisoners := CreatePrisoners()
	for _, i := range pd.PlayGames(prisoners) {
		fmt.Println(i.Results())
	}
}
