package model

type IDNResponses struct {
	Data struct {
		GetLivestreams []struct {
			Slug        string `json:"slug"`
			Title       string `json:"title"`
			ImageUrl    string `json:"image_url"`
			ViewCount   int    `json:"view_count"`
			PlaybackUrl string `json:"playback_url"`
			Status      string `json:"status"`
			LiveAt      string `json:"live_at"`
			Creator     struct {
				Username      string `json:"username"`
				Name          string `json:"name"`
				FollowerCount int    `json:"follower_count"`
			} `json:"creator"`
		} `json:"getLivestreams"`
	} `json:"data"`
}
