package main

import (
	"flag"
	"fmt"
	"log"
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

func CreatePrisoners() []pd.Prisoner {
	prisoners := []pd.Prisoner{
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

func prettyPrintMap(m map[string]int, title string) {
	// Extract keys and sort them
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Print the map in a pretty format
	fmt.Printf("%s:\n", title)
	fmt.Println("---------------")
	for _, key := range keys {
		fmt.Printf("%-15s : %d\n", key, m[key])
	}
}

// Function to calculate and print average scores
func printAverageScores(totals map[string]int, counts map[string]int, title string) {
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Printf("%s:\n", title)
	fmt.Println("---------------")
	for _, key := range keys {
		average := float64(totals[key]) / float64(counts[key])
		fmt.Printf("%-15s : %.2f\n", key, average)
	}
}

func main() {
	n := flag.Int("n", 200, "Number of games to play")
	manual := flag.Bool("manual", false, "Play the game manually via CLI")
	flag.Parse()

	if *manual {
		log.Fatalln("NO MANUAL")
		// manualPlay()
	} else {
		totals := make(map[string]int)
		oTotals := make(map[string]int)
		ownerCounts := make(map[string]int)
		prisoners := CreatePrisoners()
		for _, i := range pd.PlayGames(prisoners, *n) {
			// fmt.Println(i.Results())
			p1s, p2s := i.Scores()
			totals[i.Player1.Name] += p1s
			totals[i.Player1.Owner] += p1s
			oTotals[i.Player2.Name] += p2s
			oTotals[i.Player2.Owner] += p2s
			ownerCounts[i.Player1.Owner]++
			ownerCounts[i.Player2.Owner]++
		}
		prettyPrintMap(totals, "\nPlayer Scores")
		printAverageScores(totals, ownerCounts, "\n\nOwner Average Scores")
	}
}
