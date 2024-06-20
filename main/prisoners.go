package main

import "math/rand"


type generateChoice func(ownAnswers []bool, oppAnswers []bool) bool

type Prisoner struct {
	name     string
	owner string
	strategy generateChoice
}

func random(oa []bool, opa []bool) bool {
	return rand.Intn(2) == 1
}

func defectCooperateTwiceThenDefect(oa []bool, opa []bool) bool {
	return len(oa) % 3 != 0
}

func createPrisoners() []Prisoner {
	prisoners := []Prisoner{
		{name: "Random", owner: "Jackson", strategy: random},
		{name: "Mod 3 Prisoner", owner: "Jackson", strategy: defectCooperateTwiceThenDefect},
	}
	return prisoners
}
