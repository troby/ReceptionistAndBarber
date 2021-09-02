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
	ctx       := context.TODO()
	done      := make(chan bool, 1)

	log.Print("initializing")
	err := Init(ctx, incoming, customers, done)
	if err != nil {
		log.Fatalf("Init failed: %s", err)
	}
	log.Print("init completed")
	<-done
	log.Print("done!")
}

func Init(ctx context.Context, incoming []string, customers []*Customer, done chan bool) error {
	log.Print("show incoming")
	show(incoming)
	log.Print("show customers")
	show(customers)
	log.Print("start sleeper")
	go sleeper(ctx, done)
	return nil
}

func sleeper(ctx context.Context, done chan bool) {
	_ = ctx
	log.Printf("sleeping in sleeper")
	time.Sleep(3 * time.Second)
	log.Printf("sleeper writing to done")
	done <- true
}

func show(c interface{}) {
	fmt.Printf("%v\n", c)
}
