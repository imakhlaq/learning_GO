package main

import (
	"fmt"
	"time"
)

type server struct {
	quit    chan bool
	message chan string
}

func newServer() *server {
	return &server{
		quit:    make(chan bool),
		message: make(chan string),
	}
}

func (s *server) start() {
	fmt.Println("STARTING THE SERVER")

	//label to break loop because u cant break loop inside a select
Loop:
	for { //to keep on running the select and keep on listing for massages
		select {
		case signal := <-s.quit:
			{
				fmt.Println("closing the server", signal)
				close(s.message)
				break Loop
			}
		case message := <-s.message:
			{
				fmt.Println("Message is -> ", message)
			}
		//default is used for non blocking select if above channels have no data.
		//then don't block the goroutine
		default:
			{

			}
		}
	}
}

func req(server chan<- string) {
	server <- "HELLO MATE HOW ARE YOU"
}

func main() {

	s := newServer()
	go s.start()

	time.Sleep(3 * time.Second)
	go req(s.message)

	//only because to stop program from exiting so we can see the output
	time.Sleep(7 * time.Second)
}
