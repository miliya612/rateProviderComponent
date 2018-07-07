package rates

import (
	"math/rand"
	"time"
)

type rate struct {
	Currency string  `json:"currency"`
	Rate     float32 `json:"rate"`
}

type rates []rate

func getRates(base string, counter []string) (rates rates, err error) {
	rs, err := getRatesFromJSON()
	if err != nil {
		return nil, err
	}
	baseRate := rs.getBaseCurrencyRate(base)
	rates = []rate{}
	for _, r := range rs {
		if isReturned(counter, r, base) {
			rates = append(rates, convertBaseCurrencyOfRate(&r, baseRate))
		} else {
			continue
		}
	}
	return rates, nil
}

func convertBaseCurrencyOfRate(r *rate, base rate) rate {
	rNum := adjustRandomNumber()
	return rate{Currency: r.Currency, Rate: (r.Rate + rNum) / base.Rate}
}

func (rs rates) getBaseCurrencyRate(base string) rate {
	for _, r := range rs {
		if r.Currency == base {
			return r
		}
	}
	return rate{}
}

func adjustRandomNumber() float32 {
	rand.Seed(time.Now().UnixNano())
	return float32(rand.Intn(10)) / 10
}

func isReturned(counter []string, r rate, base string) bool {
	return contains(counter, r.Currency) && r.Currency != base
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
