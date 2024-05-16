package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavodalves/payment/internal/gateway"
)

func ProccessPayment(taxID string, value float64, gatewaies []gateway.Pay) <-chan string {
	mainCh := make(chan string)
	go func() {
		for index, gateway := range gatewaies {
			ch := make(chan string, 1)
			ctx := context.Background()

			ctx, cancel := context.WithTimeout(ctx, time.Second*4)

			defer cancel()

			go func() {
				defer close(ch)
				uuid := gateway.Pay(taxID, value)
				ch <- uuid
			}()

			select {
			case <-ctx.Done():
				{
					if index == len(gatewaies)-1 {
						panic("PAYMENT GATEWAY IS OUT")
					}
				}
			case id := <-ch:
				{
					mainCh <- id
					fmt.Println("PAYMENT PROCESSED")
					break
				}
			}
		}
	}()

	return mainCh
}

func PayUseCase(taxID string, value float64, gatewaies []gateway.Pay) {
	ch := ProccessPayment(taxID, value, gatewaies)
	fmt.Println("RESPONSE", <-ch)
}

func main() {
	gatewaies := []gateway.Pay{
		&gateway.Pay1{
			ID: byte(1),
		},
		&gateway.Pay2{
			ID: byte(2),
		},
	}

	PayUseCase("123", 50.0, gatewaies)
}
