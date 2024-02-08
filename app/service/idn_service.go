package service

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/agilistikmal/jkt48lab-htmx/app/helper"
	"github.com/agilistikmal/jkt48lab-htmx/app/model"
)

func GetIDNLives() []model.Live {
	page := 1
	var lives []model.Live
	for {
		resp := helper.GraphQLIDN(page)
		body, _ := io.ReadAll(resp.Body)
		var idnResponses model.IDNResponses
		json.Unmarshal(body, &idnResponses)
		if len(idnResponses.Data.GetLivestreams) == 0 {
			break
		}
		for _, live := range idnResponses.Data.GetLivestreams {
			if live.Status != "live" {
				continue
			}
			// if !strings.Contains(strings.ToLower(live.Slug), "jkt48") {
			// 	continue
			// }
			startedAt, _ := time.Parse("2006-01-02T15:04:05+07:00", live.LiveAt)
			lives = append(lives, model.Live{
				Member: &model.Member{
					Username:  live.Creator.Username,
					FullName:  live.Creator.Name,
					Followers: live.Creator.FollowerCount,
				},
				Title:        live.Title,
				OriginalUrl:  fmt.Sprintf("https://www.idn.app/%v/live/%v", live.Creator.Username, live.Slug),
				RoomID:       live.Slug,
				StreamingUrl: live.PlaybackUrl,
				Views:        live.ViewCount,
				Image:        live.ImageUrl,
				StartedAt:    startedAt.Unix(),
			})
		}
		page++
	}
	return lives
}
