package rates

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

var filePath = "res/rate_db.json"

func getRatesFromJSON() (rates, error) {
	var rs rates
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(f, &rs); err != nil {
		log.Fatal(err)

		return nil, err
	}
	return rs, nil
}