package src

import "fmt"

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
	Player1        Prisoner
	Player2        Prisoner
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

func WomboCombo(l []Prisoner) [][2]Prisoner {
	p2 := 1
	max := len(l)
	ret := [][2]Prisoner{}
	for p1 := 0; p1 < max && p2 <= max; {

		if p2 == max {
			p1++
			p2 = p1 + 1
			continue
		}

		combo := [2]Prisoner{l[p1], l[p2]}
		ret = append(ret, combo)
		p2++

	}
	return ret

}

func PlayGames(p []Prisoner) []Game {
	combos := WomboCombo(p)
	games := []Game{}
	for _, combo := range combos {
		player1 := combo[0]
		player2 := combo[1]
		player1choices := []bool{}
		player2choices := []bool{}
		for range gameLength {
			player1choices = append(player1choices, player1.Strategy(player1choices, player2choices))
			player2choices = append(player2choices, player2.Strategy(player2choices, player1choices))
		}
		game := Game{Player1Answers: player1choices, Player2Answers: player2choices, Player1: player1, Player2: player2}
		games = append(games, game)
	}
	return games
}
