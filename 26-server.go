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
Loop:
	for {
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
