package main

type Game struct {
	player1Answers [10]bool
	player2Answers [10]bool
	player1Points  int
	player2Points  int
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

func playGames(p *[]Prisoner) {

}
