package main

import (
	"math/rand"
	"time"
)

func clientProducer(customers chan *Client) {
	for {
		time.Sleep(time.Duration(rand.Intn(28)+7) * time.Millisecond)
		customers <- &Client{}
	}
}

func cutHear(barber *Barber, client *Client, finished chan *Barber) {
	// Cutting hear
	time.Sleep(CUTTING_TIME * time.Millisecond)
	finished <- barber
}
