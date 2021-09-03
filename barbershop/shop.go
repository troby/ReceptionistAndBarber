package barbershop

import (
	"log"
	"context"
)

func Init(ctx context.Context, incoming []string, customers chan *Customer, done chan bool) error {
	log.Printf("starting barbers")
	go startBarbers(ctx, customers, done)
	log.Print("starting receptionist")
	go startReceptionist(ctx, incoming, customers)
	return nil
}
