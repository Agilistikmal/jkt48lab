package model

type ShowroomResponses struct {
	OnLives []struct {
		GenreName string `json:"genre_name"`
		Lives     []struct {
			RoomUrlKey       string `json:"room_url_key"`
			Telop            string `json:"telop"`
			FollowerNum      int    `json:"follower_num"`
			StartedAt        int    `json:"started_at"`
			ImageSquare      string `json:"image_square"`
			ViewNum          int    `json:"view_num"`
			MainName         string `json:"main_name"`
			PremiumRoomType  int    `json:"premium_room_type"`
			RoomID           int    `json:"room_id"`
			StreamingUrlList []struct {
				Url string `json:"url"`
			} `json:"streaming_url_list"`
		} `json:"lives"`
	} `json:"onlives"`
}

type ShowroomStreamingUrlResponses struct {
	StreamingUrlList []struct {
		Label string `json:"label"`
		Url   string `json:"url"`
	} `json:"streaming_url_list"`
}
