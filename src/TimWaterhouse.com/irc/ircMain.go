package main

import (
	"fmt"
	"net"
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
	}
}
