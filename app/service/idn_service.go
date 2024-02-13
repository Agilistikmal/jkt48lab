package service

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/agilistikmal/jkt48lab/app/helper"
	"github.com/agilistikmal/jkt48lab/app/model"
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
			if !strings.Contains(live.Creator.Username, strings.ToLower(PREFIX)) {
				continue
			}
			location, _ := time.LoadLocation("Asia/Jakarta")
			startedAt, _ := time.ParseInLocation("2006-01-02T15:04:05+07:00", live.LiveAt, location)
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
				StartedAt:    fmt.Sprintf("%.2d:%.2d WIB", startedAt.Hour(), startedAt.Minute()),
			})
		}
		page++
	}
	return lives
}

func GetIDNLive(username string) model.Live {
	page := 1
	var result model.Live
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
			if !strings.Contains(live.Creator.Username, strings.ToLower(PREFIX)) {
				continue
			}
			if live.Creator.Username == username {
				startedAt, _ := time.Parse("2006-01-02T15:04:05+07:00", live.LiveAt)
				result = model.Live{
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
					StartedAt:    fmt.Sprintf("%.2d:%.2d WIB", startedAt.Hour(), startedAt.Minute()),
				}
				break
			}
		}
		page++
	}
	return result
}
