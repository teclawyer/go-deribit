package main

import (
	"fmt"
	"github.com/adampointer/go-deribit/models/private"
	flag "github.com/spf13/pflag"
	"log"

	"github.com/adampointer/go-deribit"
)

func main() {
	key := flag.String("access-key", "", "Access key")
	secret := flag.String("secret-key", "", "Secret access key")
	flag.Parse()
	errs := make(chan error)
	stop := make(chan bool)
	e, err := deribit.NewExchange(false, errs, stop)

	if err != nil {
		log.Fatalf("Error creating connection: %s", err)
	}
	if err := e.Connect(); err != nil {
		log.Fatalf("Error connecting to exchange: %s", err)
	}
	go func() {
		err := <- errs
		stop <- true
		log.Fatalf("RPC error: %s", err)
	}()
	// Hit the test RPC endpoint
	res, err := e.PublicTest(nil)
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", res.Version)
	if err := e.Authenticate(*key, *secret);err != nil {
		log.Fatalf("Error authenticating: %s", err)
	}
	summary, err := e.PrivateGetAccountSummary(&private.GetAccountSummaryRequest{Currency:"BTC"})
	if err != nil {
		log.Fatalf("Error getting account summary: %s", err)
	}
	fmt.Printf("Available funds: %f\n", summary.AvailableFunds)
	book, err := e.SubscribeBookInterval("BTC-PERPETUAL", "none", "1", "100ms")
	if err != nil {
		log.Fatalf("Error subscribing to the book: %s", err)
	}
	for b := range book {
		fmt.Printf("Top bid: %f Top ask: %f\n", b.Bids[0][0], b.Asks[0][0])
	}
}
