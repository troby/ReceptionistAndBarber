package main

import (
	"log"
	"fmt"
	"time"
	"context"
	"math/rand"
)

type Customer struct {
	Name	string
	Style	string
}

func main() {
	incoming  := []string{`bob`, `connie`, `jack`, `linda`, `paul`}
	customers := make(chan *Customer, 1)
	done      := make(chan bool, 1)

	ctx, cancel := context.WithCancel(context.TODO())
	log.Print("initializing")
	err := Init(ctx, incoming, customers, done)
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

func Init(ctx context.Context, incoming []string, customers chan *Customer, done chan bool) error {
	log.Printf("starting barbers")
	startBarbers(ctx, customers, done)
	log.Print("starting receptionist")
	go startReceptionist(ctx, incoming, customers)
	return nil
}

func startReceptionist(ctx context.Context, incoming []string, customers chan *Customer) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, name := range incoming {
				newStyle := randomStyle()
				c := new(Customer)
				c.SetName(name)
				c.SetStyle(newStyle)
				customers <- c
			}
		}
	}
}

func startBarbers(ctx context.Context, customers chan *Customer, done chan bool) {
	for i := 1; i <= 5; i++ {
		go runBarber(ctx, customers)
	}
	done <- true
}

func runBarber(ctx context.Context, customers chan *Customer) {
	for {
		select {
		case <-ctx.Done():
			return
		case c := <-customers:
			fmt.Printf("customer: %s, style: %s\n", c.Name, c.Style)
		}
	}
}

func randomStyle() string {
	styles := []string{`cut`, `trim`, `perm`}
	index := rand.Intn(3)
	//log.Printf("random index: %d", index)
	return styles[index]
}

func (c *Customer) SetName(name string) {
	c.Name = name
}

func (c *Customer) SetStyle(style string) {
	c.Style = style
}
