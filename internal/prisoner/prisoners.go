package prisoner

type generateChoice func(ownAnswers []bool, oppAnswers []bool) bool

type Prisoner struct {
	Name     string
	Owner    string
	Strategy generateChoice
}
