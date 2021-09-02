/*
 Initialize:
   - create []incoming : slice of new arrivals (possible customers)
   - create []customers : slice of customers waiting for a haircut
   - create Receptionist() : lifetime thread that creates customers out of incoming data
   - create Barbers() : lifetime threads that check for customers and performs Validate, Ownership, and Provision
*/

package main

import (
	"log"
	"fmt"
	"sync"
	"time"
	"context"
)

type Customer struct {
	Name	string
	Type	string
	Wait	sync.Mutex
}

func main() {
	incoming  := []string{}
	customers := []*Customer{}

	ctx, cancel := context.WithCancel(context.TODO())
	log.Print("initializing")
	err := Init(ctx, incoming, customers)
	if err != nil {
		log.Fatalf("Init failed: %s", err)
	}
	log.Print("init completed")
	time.Sleep(20 * time.Second)
	log.Print("cancelling sleeps")
	cancel()
	log.Print("done!")
}

func Init(ctx context.Context, incoming []string, customers []*Customer) error {
	log.Print("show incoming")
	show(incoming)
	log.Print("show customers")
	show(customers)
	log.Print("start sleeper")
	go sleeper(ctx)
	return nil
}

func sleeper(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Printf("sleeping in sleeper %d times.", n)
			time.Sleep(3 * time.Second)
			n += 1
		}
	}
}

func show(c interface{}) {
	fmt.Printf("%v\n", c)
}
