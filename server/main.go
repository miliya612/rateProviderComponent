package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/miliya612/crpc/rateservice"
	"google.golang.org/grpc"
	"context"
	cs "github.com/miliya612/crpc/server/currencies"
	rs "github.com/miliya612/crpc/server/rates"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type rateProviderServer struct {
}

func (s *rateProviderServer) GetRate(ctx context.Context, req *pb.RateRequest) (res *pb.RateResponse, err error) {
	fmt.Println(rs.GetRate(req))
	return rs.GetRate(req)
}

func (s *rateProviderServer) GetSupportedCurrencies(ctx context.Context, e *pb.Empty) (res *pb.Currencies, err error) {
	return cs.GetSupportedCurrencies()
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRateProviderServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func newServer() *rateProviderServer {
	s := &rateProviderServer{}
	return s
}
