package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"

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

func prettyPrintMap(m map[string]int) {
	// Extract keys and sort them
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print the map in a pretty format
	fmt.Println("Player Scores:")
	fmt.Println("---------------")
	for _, key := range keys {
		fmt.Printf("%-15s : %d\n", key, m[key])
	}
}

func main() {
	n := flag.Int("n", 200, "Number of games to play")
	manual := flag.Bool("manual", false, "Play the game manually via CLI")
	flag.Parse()

	if *manual {
		panic("NOT IMPLEMENTED")
		// manualPlay()
	} else {
		totals := make(map[string]int) // Initialize the map
		prisoners := CreatePrisoners()
		for _, i := range pd.PlayGames(prisoners, *n) {
			fmt.Println(i.Results())
			p1s, p2s := i.Scores()
			totals[i.Player1.Name] += p1s
			totals[i.Player2.Name] += p2s
		}
		prettyPrintMap(totals)
	}
}
