package main

import "time"

func main() {
	customers := make(chan *Client)

	go clientProducer(customers)
	go BarberShop(customers)
	time.Sleep(2 * time.Second)
}
