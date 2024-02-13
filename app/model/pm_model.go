package model

type PMStats struct {
	Type   string
	Period struct {
		From string
		To   string
	}
	Ranks []struct {
		Name       string
		Points     int
		Count      int
		TextCount  int
		ImageCount int
		VoiceCount int
	}
}
