package main

import (
	"fmt"
)

func BarberShop(customers <-chan *Client) {
	freeBarbers := []*Barber{}
	waitingClient := []*Client{}
	syncBarberChan := make(chan *Barber)

	//creating barbers
	for i := 0; i < BARBERS_AMOUNT; i++ {
		freeBarbers = append(freeBarbers, &Barber{})
	}

	for {
		select {
		case client := <-customers:
			if len(freeBarbers) == 0 {
				if len(waitingClient) < HALL_SITS_AMOUNT {
					// client is waiting in the hall
					waitingClient = append(waitingClient, client)
					fmt.Printf("Client is waiting in hall (%v)\n", len(waitingClient))
				} else {
					// hall is full - bye-bye client
					fmt.Println("No free space for client")
				}
			} else {
				barber := freeBarbers[0]
				freeBarbers = freeBarbers[1:]

				fmt.Println("Client goes to barber for haircut")
				go cutHear(barber, client, syncBarberChan)
			}
		// barber finish work
		case barber := <-syncBarberChan:
			if len(waitingClient) > 0 {
				// get client from hall
				client := waitingClient[0]
				waitingClient = waitingClient[1:]

				fmt.Printf("Take client from room (%v)\n", len(waitingClient))
				go cutHear(barber, client, syncBarberChan)
			} else {
				// barber is going to sleep
				fmt.Println("Barber goes to Zzz...")
				freeBarbers = append(freeBarbers, barber)
			}
		}
	}
}
