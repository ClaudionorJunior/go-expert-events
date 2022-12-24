package main

import (
	"strconv"

	"github.com/ClaudionorJunior/go-expert-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 1000000; i++ {
		intConverted := strconv.Itoa(i)
		rabbitmq.Publish(ch, "Hello World!"+intConverted, "amq.direct")

	}

}
