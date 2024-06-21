package main

import (
	"fmt"
	"math/rand"

	pd "github.com/jcpage573/prisoners-dilemma/src"
)

// Example Strategies

func random(oa []bool, opa []bool) bool {
	return rand.Intn(2) == 1
}

func defectCooperateTwiceThenDefect(oa []bool, opa []bool) bool {
	return len(oa)%3 != 0
}

func CreatePrisoners() []pd.Prisoner {
	prisoners := []pd.Prisoner{
		{Name: "Random", Owner: "Jackson", Strategy: random},
		{Name: "Mod 3 Prisoner", Owner: "Jackson", Strategy: defectCooperateTwiceThenDefect},
	}
	return prisoners
}

func main() {
	prisoners := CreatePrisoners()
	for _, i := range pd.PlayGames(prisoners) {
		fmt.Println(i.Results())
	}
}
