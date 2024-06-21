package main

import (
	"flag"
	"log"
	"math/rand"

	"github.com/jcpage573/prisoners-dilemma/internal/game"
	"github.com/jcpage573/prisoners-dilemma/internal/prisoner"
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

func titForTatWithForgiveness(oa []bool, opa []bool) bool {
	n := len(oa)
	if n < 1 {
		return true
	}
	// 10% chance to forgive a defection
	if !opa[n-1] && rand.Float64() < 0.1 {
		return true
	}
	return opa[n-1]
}

func megatron(oa []bool, opa []bool) bool {
	if len(oa) < 10 {
		return rand.Intn(2) == 1
	}

	for _, a := range opa {
		if !a {
			return false
		}
	}
	return rand.Intn(2) == 1
}

func CreatePrisoners() []prisoner.Prisoner {
	prisoners := []prisoner.Prisoner{
		// Jackson
		{Name: "Random", Owner: "Jackson", Strategy: random},
		{Name: "Mod 3 Prisoner", Owner: "Jackson", Strategy: defectCooperateTwiceThenDefect},
		{Name: "Defect", Owner: "Jackson", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "Jackson", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "Jackson", Strategy: func([]bool, []bool) bool { return false }},

		// Hunter
		{Name: "Tit for Tat", Owner: "Hunter", Strategy: titForTat},
		{Name: "Megaton", Owner: "Hunter", Strategy: megatron},
		{Name: "Defect", Owner: "Hunter", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "Hunter", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "Hunter", Strategy: func([]bool, []bool) bool { return false }},

		// ChatGPT
		{Name: "T4T Forgiveness", Owner: "ChatGPT", Strategy: titForTatWithForgiveness}, // Alex
		{Name: "Defect", Owner: "ChatGPT", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "ChatGPT", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "ChatGPT", Strategy: func([]bool, []bool) bool { return false }},
		{Name: "Defect", Owner: "ChatGPT", Strategy: func([]bool, []bool) bool { return false }},
	}
	return prisoners
}

func main() {
	n := flag.Int("n", 200, "Number of games to play")
	manual := flag.Bool("manual", false, "Play the game manually via CLI")
	flag.Parse()

	if *manual {
		log.Fatalln("NO MANUAL")
		// manualPlay()
	} else {
		prisoners := CreatePrisoners()
		_, totals := game.Play(prisoners, *n)
		totals.PrintTotals()
	}
}
