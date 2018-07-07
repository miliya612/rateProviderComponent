package rates

import (
	"math/rand"
	"time"
)

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
	for i, r := range rs {
		rs[i] = convertBaseCurrencyOfRate(&r, baseRate)
	}
	return rs, nil
}

func convertBaseCurrencyOfRate(r *rate, base rate) rate {
	rNum := adjustRandomNumber()
	return rate{Currency:r.Currency, Rate:(r.Rate+rNum)/base.Rate}
}

func (rs rates)getBaseCurrencyRate(base string) rate {
	for _, r := range rs {
		if r.Currency == base {
			return r
		}
	}
	return rate{}
}

func adjustRandomNumber() float32 {
	rand.Seed(time.Now().UnixNano())
	return float32(rand.Intn(10))/10
}
