package model

type ShowroomResponses struct {
	OnLives []struct {
		Lives []struct {
			RoomUrlKey      string `json:"room_url_key"`
			Telop           string `json:"telop"`
			FollowerNum     int    `json:"follower_num"`
			StartedAt       int    `json:"started_at"`
			Image           string `json:"image"`
			ViewNum         int    `json:"view_num"`
			MainName        string `json:"main_name"`
			PremiumRoomType int    `json:"premium_room_type"`
			RoomID          int    `json:"room_id"`
		} `json:"lives"`
	} `json:"onlives"`
}
