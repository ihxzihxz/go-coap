package main

import (
	"log"
	"os"

	"github.com/ihxzihxz/go-coap"
)

func main() {

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 12345,
		Payload:   []byte("hello, world!"),
	}

	path := "/a"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	req.SetOption(coap.ETag, "weetag")
	req.SetOption(coap.MaxAge, 3)
	req.SetPathString(path)

	c, err := coap.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		log.Printf("Response payload: %s", rv.Payload)
	}

}
