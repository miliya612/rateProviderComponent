package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	pb "github.com/miliya612/crpc/rateservice"
	"context"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRateProviderClient(conn)
	printRates(client, &pb.RateRequest{Base: "USD", Counter: []string{"JPY", "TWD", "USD"}})
	printCurrencies(client, &pb.Empty{})
}

// printRates gets the rates respond to the given base currency.
func printRates(client pb.RateProviderClient, req *pb.RateRequest) {
	log.Printf("Getting releated rate for requets (BASE: %s, COUNTER: %s)", req.Base, req.Counter)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rates, err := client.GetRate(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetRate(_) = _, %v: ", client, err)
	}
	log.Printf("%v", rates)
}

// loadCurrencies gets the rates respond to the given base currency.
func printCurrencies(client pb.RateProviderClient, empty *pb.Empty) {
	log.Println("Getting supported currencies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rates, err := client.GetSupportedCurrencies(ctx, empty)
	if err != nil {
		log.Fatalf("%v.GetSupportedCurrencies(_) = _, %v: ", client, err)
	}
	log.Printf("%v", rates)
}