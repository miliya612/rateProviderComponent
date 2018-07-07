package currencies

import (
	pb "github.com/miliya612/crpc/rateservice"
)

func GetSupportedCurrencies() (res *pb.Currencies, err error) {
	cs, err := getCurrencies()
	if err != nil {
		return nil, err
	}
	p := pb.Currencies{}
	return cs.translate(&p), nil
}

func (cs currencies)translate(to *pb.Currencies) *pb.Currencies {
	if len(cs) == 0 {
		return to
	}
	for _, cur := range cs {
		to.Currency = append(to.Currency, cur.Currency)
	}
	return to
}