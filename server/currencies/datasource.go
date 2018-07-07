package currencies

import (
	"io/ioutil"
	"encoding/json"
)

var filePath = "res/currencies_db.json"

func getCurrenciesFromJSON() (currencies, error) {
	var cs currencies
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, &cs); err != nil {
		return nil, err
	}
	return cs, nil
}