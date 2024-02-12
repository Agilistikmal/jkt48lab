package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/agilistikmal/jkt48lab/app/model"
)

func GetSRLives() []model.Live {
	var lives []model.Live
	resp, _ := http.Get("https://www.showroom-live.com/api/live/onlives")
	body, _ := io.ReadAll(resp.Body)
	var srResponses model.ShowroomResponses
	json.Unmarshal(body, &srResponses)
	if len(srResponses.OnLives[0].Lives) == 0 {
		return nil
	}
	for _, genre := range srResponses.OnLives {
		if genre.GenreName != "Idol" {
			continue
		}
		for _, live := range genre.Lives {
			if !strings.Contains(live.RoomUrlKey, PREFIX) {
				continue
			}
			if live.PremiumRoomType != 0 {
				continue
			}
			streamingUrl := GetSRStreamingUrl(live.RoomID)
			location, _ := time.LoadLocation("Asia/Jakarta")
			startedAt := time.Unix(int64(live.StartedAt), 0).In(location)
			lives = append(lives, model.Live{
				Member: &model.Member{
					Username:  live.RoomUrlKey,
					FullName:  live.MainName,
					Followers: live.FollowerNum,
				},
				Title:        live.Telop,
				RoomID:       string(rune(live.RoomID)),
				OriginalUrl:  fmt.Sprintf("https://showroom-live.com/r/%v", live.RoomUrlKey),
				StreamingUrl: streamingUrl,
				Views:        live.ViewNum,
				Image:        live.ImageSquare,
				StartedAt:    fmt.Sprintf("%.2d:%.2d WIB", startedAt.Hour(), startedAt.Minute()),
			})
		}
	}
	return lives
}

func GetSRLive(username string) model.Live {
	lives := GetSRLives()
	var result model.Live
	for _, live := range lives {
		if live.Member.Username == username {
			result = live
			break
		}
	}
	return result
}

func GetSRStreamingUrl(roomID int) string {
	resp, _ := http.Get(fmt.Sprintf("https://www.showroom-live.com/api/live/streaming_url?abr_available=1&room_id=%v", roomID))
	body, _ := io.ReadAll(resp.Body)
	var srStreamingUrlResponses model.ShowroomStreamingUrlResponses
	json.Unmarshal(body, &srStreamingUrlResponses)
	if len(srStreamingUrlResponses.StreamingUrlList) == 0 {
		return ""
	}
	return srStreamingUrlResponses.StreamingUrlList[1].Url
}
