package barbershop

import (
	"fmt"
	"sync"
	"context"
)

func startBarbers(ctx context.Context, customers chan *Customer, done chan bool) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go runBarber(ctx, customers, &wg)
	}
	wg.Wait()
	done <- true
}

func runBarber(ctx context.Context, customers chan *Customer, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case c := <-customers:
			fmt.Printf("customer: %s, style: %s\n", c.Name, c.Style)
		}
	}
}
