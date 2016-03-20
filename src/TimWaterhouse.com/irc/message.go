package main

import (
	"bytes"
	"container/list"
	"fmt"
	"net"
	"strings"
)

type Message struct {
	prefix  string
	command string
	params  []string
}

func ReadUntilSpace(conn net.Conn) (string, bool) {
	var messageBuffer bytes.Buffer
	var eol bool = false

	for {
		var buffer = make([]byte, 1)
		var _, _ = conn.Read(buffer)

		if buffer[0] == ' ' {
			break
		} else if buffer[0] == '\n' {
			eol = true
			break
		} else if buffer[0] == '\r' {
			continue
		}

		messageBuffer.Write(buffer)
	}

	return messageBuffer.String(), eol
}

func ReadUntilSpaceOrEOL(conn net.Conn) (string, bool) {
	var messageBuffer bytes.Buffer
	var eol bool = false
	var firstChar bool = true
	var readToEOL = false

	for {
		var buffer = make([]byte, 1)
		var _, _ = conn.Read(buffer)

		if firstChar && buffer[0] == ':' {
			readToEOL = true
		}
		firstChar = false

		if !readToEOL && buffer[0] == ' ' {
			break
		} else if buffer[0] == '\n' {
			eol = true
			break
		} else if buffer[0] == '\r' {
			continue
		}

		messageBuffer.Write(buffer)
	}

	return messageBuffer.String(), eol
}

/*
<message>  ::= [':' <prefix> <SPACE> ] <command> <params> <crlf>
<prefix>   ::= <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
<command>  ::= <letter> { <letter> } | <number> <number> <number>
<SPACE>    ::= ' ' { ' ' }
<params>   ::= <SPACE> [ ':' <trailing> | <middle> <params> ]

<middle>   ::= <Any *non-empty* sequence of octets not including SPACE
               or NUL or CR or LF, the first of which may not be ':'>
<trailing> ::= <Any, possibly *empty*, sequence of octets not including
                 NUL or CR or LF>

<crlf>     ::= CR LF
*/
func ReadMessage(conn net.Conn) Message {
	var m Message
	var eol bool = false

	var field string
	field, eol = ReadUntilSpace(conn)
	if field[0] == ':' { // Indicates this is a prefix
		m.prefix = field

		if eol {
			// TODO: Error here since we shouldn't get just a prefix
			return m
		}

		field, eol = ReadUntilSpace(conn)
	}

	m.command = field

	params := list.New()
	for !eol {
		field, eol = ReadUntilSpaceOrEOL(conn)
		params.PushBack(field)
	}

	paramsList := make([]string, params.Len())
	for p := params.Front(); p != nil; p = p.Next() {
		paramsList = append(paramsList, p.Value.(string))
	}

	m.params = paramsList

	return m
}

func (m *Message) ToString() string {
	var output string
	if m.prefix != "" {
		output += m.prefix + " "
	}

	if m.command != "" {
		output += m.command + " "
	}

	return output + strings.Join(m.params, " ")
}

func (m *Message) Send(conn net.Conn) {
	fmt.Fprintf(conn, m.ToString()+"\r\n")
}
