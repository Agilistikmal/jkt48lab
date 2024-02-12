package model

type Live struct {
	Member       *Member
	Title        string
	RoomID       string
	OriginalUrl  string
	StreamingUrl string
	Views        int
	Image        string
	StartedAt    string
}
