package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GraphQLIDN(page int) *http.Response {
	query, err := json.Marshal(map[string]any{
		"query": fmt.Sprintf(`
			query GetLivestreams {
				getLivestreams(page: %v) {
					slug
					title
					image_url
					view_count
					playback_url
					status
					live_at
					gift_icon_url
					creator {
							username
							name
							follower_count
					}
				}
			}
    	`, page),
	})
	gReq, _ := http.NewRequest("POST", "https://api.idn.app/graphql", bytes.NewBuffer(query))
	gReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	var resp *http.Response
	for {
		resp, err = client.Do(gReq)
		if err == nil {
			break
		}
	}
	return resp
}
