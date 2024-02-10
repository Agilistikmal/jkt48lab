package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/agilistikmal/jkt48lab-htmx/app/model"
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
			lives = append(lives, model.Live{
				Member: &model.Member{
					Username:  live.RoomUrlKey,
					FullName:  live.MainName,
					Followers: live.FollowerNum,
				},
				Title:        live.Telop,
				RoomID:       string(rune(live.RoomID)),
				OriginalUrl:  fmt.Sprintf("https://showroom-live.com/r/%v", live.RoomUrlKey),
				StreamingUrl: live.StreamingUrlList[0].Url,
				Views:        live.ViewNum,
				Image:        live.Image,
				StartedAt:    int64(live.StartedAt),
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
