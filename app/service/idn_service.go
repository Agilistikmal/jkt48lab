package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
			if !strings.Contains(strings.ToUpper(live.Slug), PREFIX) {
				continue
			}
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

func GetIDNLive(username string) model.Live {
	page := 1
	var result model.Live
	for {
		resp := helper.GraphQLIDN(page)
		defer resp.Body.Close()
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
			if !strings.Contains(strings.ToLower(live.Slug), strings.ToLower(PREFIX)) {
				continue
			}
			if live.Creator.Username == username {
				var streamingUrl string
				startedAt, _ := time.Parse("2006-01-02T15:04:05+07:00", live.LiveAt)

				playbackResp, err := http.Get(live.PlaybackUrl)
				if err != nil {
					log.Fatal(err)
				}
				defer playbackResp.Body.Close()
				body, _ := io.ReadAll(playbackResp.Body)
				content := strings.Split(string(body), "\n")
				for _, c := range content {
					if strings.Contains(c, ".m3u8") {
						streamingUrl = c
					}
				}

				result = model.Live{
					Member: &model.Member{
						Username:  live.Creator.Username,
						FullName:  live.Creator.Name,
						Followers: live.Creator.FollowerCount,
					},
					Title:        live.Title,
					OriginalUrl:  fmt.Sprintf("https://www.idn.app/%v/live/%v", live.Creator.Username, live.Slug),
					RoomID:       live.Slug,
					StreamingUrl: streamingUrl,
					Views:        live.ViewCount,
					Image:        live.ImageUrl,
					StartedAt:    startedAt.Unix(),
				}
				break
			}
		}
		page++
	}
	return result
}
