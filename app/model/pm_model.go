package model

type PMStats struct {
	Type   string
	Period struct {
		From string
		To   string
	}
	Ranks []struct {
		Member *Member
		Points int
		Total  int
		Text   int
		Image  int
		Voice  int
	}
}
