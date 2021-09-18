package main

import (
	"flag"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

var sub, pub *string

func main() {
	sub = flag.String("sub", "127.0.0.1:9000", "address for subscribing")
	pub = flag.String("pub", "100.200.80.1:9000", "address for publishing")
	flag.Parse()

	// establishing publisher
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	err := publisher.Bind("tcp://" + *pub)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publisher Binded Successfully...", *pub)
	}

	// establishing subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	err = subscriber.Connect("tcp://" + *sub)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Subscriber Connected Successfully...", *sub)
	}
	subscriber.SetSubscribe("")

	for {
		contents, _ := subscriber.RecvBytes(0)

		publisher.SendBytes(contents, 0)
	}
}
