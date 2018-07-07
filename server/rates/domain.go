package rates

type rate struct {
	Currency string `json:"currency"`
	Rate float32 `json:"rate"`
}

type rates []rate

func getRates(base string) (rs rates, err error){
	rs, err = getRatesFromJSON()
	if err != nil {
		return nil, err
	}
	baseRate := rs.getBaseCurrencyRate(base)
	for _, r := range rs {
		r = convertBaseCurrencyOfRate(r, baseRate)
	}
	return rs, nil
}

func convertBaseCurrencyOfRate(r rate, base rate) rate {
	r.Rate = r.Rate/base.Rate
	return r
}

func (rs rates)getBaseCurrencyRate(base string) rate {
	for _, r := range rs {
		if r.Currency == base {
			return r
		}
	}
	return rate{}
}
