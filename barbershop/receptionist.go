package barbershop

import (
	"context"
	"math/rand"
)

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

func randomStyle() string {
	styles := []string{`cut`, `trim`, `perm`}
	index := rand.Intn(3)
	return styles[index]
}
