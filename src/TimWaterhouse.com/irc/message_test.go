package irc

import "testing"

func TestMessageToString(t *testing.T) {
	var m Message
	m.prefix = "prefix"
	m.command = "command"
	m.params = make([]string, 3)
	m.params[0] = "1"
	m.params[1] = "2"
	m.params[2] = "3"

	var expected = "prefix command 1 2 3"
	var got = m.ToString()
	if got != expected {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}
