package game

import (
	"fmt"
	"sort"

	"github.com/jcpage573/prisoners-dilemma/internal/prisoner"
)

const (
	bothTrue   = 3
	sucker     = 0
	thief      = 5
	bothFalse  = 1
	gameLength = 3
)

type Game struct {
	Player1Answers []bool
	Player2Answers []bool
	Player1        prisoner.Prisoner
	Player2        prisoner.Prisoner
}

func Sim(prisoners []prisoner.Prisoner, n int) {
	totals := make(map[string]int)
	oTotals := make(map[string]int)
	ownerCounts := make(map[string]int)
	for _, i := range Play(prisoners, n) {
		// fmt.Println(i.Results())
		p1s, p2s := i.Scores()
		totals[i.Player1.Name] += p1s
		oTotals[i.Player1.Owner] += p1s
		totals[i.Player2.Name] += p2s
		oTotals[i.Player2.Owner] += p2s
		ownerCounts[i.Player1.Owner]++
		ownerCounts[i.Player2.Owner]++
	}
	prettyPrintMap(totals, "\nPlayer Scores")
	printAverageScores(oTotals, ownerCounts, "\n\nOwner Average Scores")
}

func (g Game) Scores() (p1 int, p2 int) {
	p1Score := 0
	p2Score := 0
	for i := range g.Player1Answers {
		player1choice := g.Player1Answers[i]
		player2choice := g.Player2Answers[i]
		if player1choice && player2choice {
			p1Score += bothTrue
			p2Score += bothTrue
		} else if !(player1choice || player2choice) {
			p1Score += bothFalse
			p2Score += bothFalse
		} else if player1choice {
			p2Score += thief
		} else if player2choice {
			p1Score += thief
		}
	}
	return p1Score, p2Score
}

func (g Game) Results() string {
	p1score, p2score := g.Scores()
	res := fmt.Sprintf("%-15s %-10s\n", "Prisoner", "Score")
	res += fmt.Sprintf("%-15s %-10d\n", g.Player1.Name, p1score)
	res += fmt.Sprintf("%-15s %-10d\n", g.Player2.Name, p2score)
	return res
}

func Play(p []prisoner.Prisoner, n int) []Game {
	games := []Game{}
	for i := 0; i < n; i++ {
		for _, combo := range WomboCombo(p) {
			player1 := combo[0]
			player2 := combo[1]
			player1choices := []bool{}
			player2choices := []bool{}
			for j := 0; j < gameLength; j++ {
				player1choices = append(player1choices, player1.Strategy(player1choices, player2choices))
				player2choices = append(player2choices, player2.Strategy(player2choices, player1choices))
			}
			game := Game{Player1Answers: player1choices, Player2Answers: player2choices, Player1: player1, Player2: player2}
			games = append(games, game)
		}
	}
	return games
}

func WomboCombo(l []prisoner.Prisoner) [][2]prisoner.Prisoner {
	p2 := 1
	max := len(l)
	ret := [][2]prisoner.Prisoner{}
	for p1 := 0; p1 < max && p2 <= max; {

		if p2 == max {
			p1++
			p2 = p1 + 1
			continue
		}

		combo := [2]prisoner.Prisoner{l[p1], l[p2]}
		ret = append(ret, combo)
		p2++

	}
	return ret
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
