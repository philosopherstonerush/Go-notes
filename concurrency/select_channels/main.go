package main

import "fmt"

type server struct {
	quitch chan struct{} // empty struct
	msgch  chan string   // a message channel
}

func (s *server) start() {
	fmt.Println("Server is starting...")
	s.loop()
}

func (s *server) loop() {
mainloop:
	for {
		select { // for - select loop, used for iterating over the channels, randomly selects them, their order doesnt matter
		case <-s.quitch: // empty struct so it doesnt proceed unless we inilise it with something
			break mainloop // breaks the for loop, without it, only select loop gets broken
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default: // shouldnt use this irl but without it this wont run
		}
	}
}

func (s *server) handleMessage(msg string) {
	fmt.Printf("The message is %v", msg)
}

func main() {
	s := &server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
	go s.start()
	s.msgch <- "HEllO"
}
