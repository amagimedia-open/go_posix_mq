package main

import (
	"fmt"
	"log"
	"time"

	posix_mq "github.com/amagimedia-open/go_posix_mq"
)

const maxSendTickNum = 10

func main() {
	oflag := posix_mq.O_WRONLY | posix_mq.O_CREAT
	mq, err := posix_mq.NewMessageQueue("/posix_mq_example", oflag, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(mq *posix_mq.MessageQueue) {
		err := mq.Close()
		if err != nil {
			log.Println(err)
		}
	}(mq)

	count := 0
	for {
		count++
		err = mq.Send([]byte(fmt.Sprintf("Hello, World : %d\n", count)), 0)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Sent a new message")

		if count >= maxSendTickNum {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
