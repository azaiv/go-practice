package _select

import (
	"fmt"
	"strconv"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func CreateMessage() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Друг 1",
				Text:   "Привет",
			}

			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Друг 2",
				Text:   "Как дела?",
			}

			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-messageChan1:
			fmt.Println("Я получил сообщение от: ", msg1.Author, "текст сообщений: ", msg1.Text)
		case msg2 := <-messageChan2:
			fmt.Println("Я получил сообщение от: ", msg2.Author, "текст сообщений: ", msg2.Text)
		}
	}
}

func Test() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case number := <-intCh:
			fmt.Println("intCh", number)
		case str := <-strCh:
			fmt.Println("strCh", str)
		}
	}
}
