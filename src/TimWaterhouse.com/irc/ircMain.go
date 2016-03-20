package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func drawFrame() {
	termbox.SetCell(0, 0, 'A', termbox.ColorGreen, termbox.ColorBlack)
	printf_tb(3, 18, termbox.ColorWhite, termbox.ColorBlack, "Foooooo")
}

func main() {
	tbErr := termbox.Init()

	if tbErr != nil {
		panic(tbErr)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawFrame()
	termbox.Flush()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			}
		}
	}

	/*fmt.Printf("Welcome to irc, enter server name with port: ")
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

	for {
		var message = ReadMessage(conn)
		fmt.Print(message.ToString() + "\n")

		if message.command == "PING" {
			var response Message
			response.command = "PONG"
			response.params = message.params
			response.Send(conn)
			fmt.Print(response.ToString() + "\n")
		}
	}*/
}
