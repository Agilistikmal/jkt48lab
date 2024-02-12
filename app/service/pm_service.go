package service

import (
	"encoding/json"
	"log"

	"github.com/agilistikmal/jkt48lab/app/db"
	"github.com/agilistikmal/jkt48lab/app/model"
)

func GetPMStats(statsType string) []model.PMStats {
	conn := db.OpenConnection()
	defer conn.Close()

	rows, err := conn.Query(`
		SELECT data FROM pm_stats WHERE 
		data->>'type' = $1
	`, statsType)

	if err != nil {
		log.Fatal(err)
	}

	var ranks []model.PMStats
	for rows.Next() {
		var res string
		var pmStats model.PMStats
		rows.Scan(&res)
		json.Unmarshal([]byte(res), &pmStats)
		ranks = append(ranks, pmStats)
	}

	return ranks
}
