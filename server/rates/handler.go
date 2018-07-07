package rates

import (
	pb "github.com/miliya612/crpc/rateservice"
	"time"
)

func GetRate(req *pb.RateRequest) (res *pb.RateResponse, err error) {
	rs, err := getRates(req.Base, req.Counter)
	if err != nil {
		return nil, err
	}
	res = &pb.RateResponse{}
	res.Base = req.Base
	res.Date = timeToDateString(time.Now())
	res.Rates = rs.translate([]*pb.Rates{})
	return res, nil
}

func (rs rates)translate(to []*pb.Rates) []*pb.Rates {
	if len(rs) == 0 {
		return to
	}
	for _, r := range rs {
		to = append(to, &pb.Rates{Currency:r.Currency, Value:r.Rate})
	}
	return to
}

func timeToDateString(t time.Time) string {
	// refer http://text.baldanders.info/golang/time-functions/
	const layout = "2006-01-02"
	return t.Format(layout)
}