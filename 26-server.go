package main

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
	for {
		select {
		case signal <- s.quit:
			{

			}

		}
	}
}
func main() {

	s := newServer()

	s.start()

}
