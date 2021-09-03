package main

import (
	"log"
	"time"
	"context"
	"ReceptionistAndBarber/barbershop"
)

func main() {
	incoming  := []string{`bob`, `connie`, `jack`, `linda`, `paul`}
	customers := make(chan *barbershop.Customer, 1)
	done      := make(chan bool, 1)

	ctx, cancel := context.WithCancel(context.TODO())
	log.Print("initializing")
	err := barbershop.Init(ctx, incoming, customers, done)
	if err != nil {
		log.Fatalf("Init failed: %s", err)
	}
	log.Print("init completed")

	time.Sleep(5 * time.Second)
	log.Print("canceling tasks")
	cancel()

	<-done
	log.Print("Done!")
}
