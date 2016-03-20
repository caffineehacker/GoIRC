package main

import (
	"container/list"

	"github.com/nsf/termbox-go"
	"timwaterhouse.com/irc/irc"
)

type View struct {
	inputString string
	messages    list.List
}

func (v *View) drawMessages() {
	var width, height = termbox.Size()
	var y = height - (3 + (len(v.inputString) / width))
	for msg := v.messages.Back(); msg != nil; msg = msg.Prev() {
		msgStr := msg.Value.(irc.Message).ToString()
		var linesRequired = len(msgStr)/width + 1

		for ; y >= 0 && linesRequired > 0; y-- {
			var x int
			for charIndex := width * (linesRequired - 1); charIndex < len(msgStr) && x < width; charIndex++ {
				termbox.SetCell(x, y, []rune(msgStr)[charIndex], termbox.ColorWhite, termbox.ColorBlack)
				x++
			}

			linesRequired--
		}
	}
}

func (v *View) AppendMessage(message irc.Message) {
	v.messages.PushBack(message)
	v.drawIrc()
}

func (v *View) SetInputString(input string) {
	v.inputString = input
	v.drawIrc()
}

func (v *View) drawIrc() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	v.drawInputString()
	v.drawMessages()
	termbox.Flush()
}

func (v *View) drawInputString() {
	var width, height = termbox.Size()
	var linesRequired = len(v.inputString)/width + 1

	var charIndex int
	for y := height - (1 + linesRequired); y >= 0 && y < height-1; y++ {
		for x := 0; x < width; x++ {
			var c rune
			if charIndex < len(v.inputString) {
				c = []rune(v.inputString)[charIndex]
			} else {
				c = ' '
			}
			termbox.SetCell(x, y, c, termbox.ColorBlack, termbox.ColorGreen)

			charIndex++
		}
	}
}
