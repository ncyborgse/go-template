package helloworld

import (
	"errors"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

type HelloWorld struct {
	msg string
}
var counter int
var mu sync.Mutex

func NewHelloWorld() *HelloWorld {
	err := errors.New("This is an error")

	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("Error detected")
	}

	go increment()
	go increment()

	return &HelloWorld{
		msg: "Hello, World!",
	}
}

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func (hello *HelloWorld) Talk() {
	log.WithFields(log.Fields{"Msg": hello.msg, "OtherMsg": "Logging is cool!"}).Info("Talking...")
	fmt.Println(hello.msg)
}
