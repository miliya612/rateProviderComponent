package currencies

type currencies []currency

type currency struct {
	Currency string `json:"currency"`
}

func getCurrencies() (currencies currencies, err error){
	return getCurrenciesFromJSON()
}
