package barbershop

import (
	"fmt"
	"context"
)

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
