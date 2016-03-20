package main

import (
	"container/list"
	"fmt"
	"net"

	"github.com/nsf/termbox-go"
	"timwaterhouse.com/irc/irc"
)

func main() {
	fmt.Printf("Welcome to irc, enter server name with port: ")
	var serverName string
	fmt.Scanf("%s\n", &serverName)

	var conn, err = net.Dial("tcp", serverName)
	if err != nil {
		return
	}

	fmt.Printf("Enter nickname: ")
	var nickName string
	fmt.Scanf("%s\n", &nickName)
	fmt.Fprintf(conn, "NICK "+nickName+"\r\n")

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanf("%s\n", &username)
	fmt.Printf("Enter real name: ")
	var realName string
	fmt.Scanf("%s\n", &realName)
	fmt.Fprintf(conn, "USER "+username+" 0 * :"+realName+"\r\n")

	// Start termbox
	tbErr := termbox.Init()

	if tbErr != nil {
		panic(tbErr)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	messageTexts := list.New()
	drawIrc(messageTexts)

	messageCh := make(chan string, 5)

	go readMessages(conn, messageCh)

	var inputString string
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyEnter {
				// TODO: Handle entering a command
				inputString = ""
			} else {
				inputString += string(ev.Ch)
			}

			drawInputString(inputString)
			break
		case termbox.EventInterrupt:
			msg := <-messageCh
			messageTexts.PushBack(msg)
			drawIrc(messageTexts)
			break
		}
	}
}

func readMessages(conn net.Conn, messageCh chan string) {
	for {
		var message = irc.ReadMessage(conn)

		if message.Command == "PING" {
			var response irc.Message
			response.Command = "PONG"
			response.Params = message.Params
			response.Send(conn)
			fmt.Print(response.ToString() + "\n")
		} else {
			messageCh <- message.ToString() + "\n"
			termbox.Interrupt()
		}
	}
}
