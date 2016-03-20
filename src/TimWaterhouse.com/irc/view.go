package main

import (
	"container/list"

	"github.com/nsf/termbox-go"
)

//text string

func drawSeparator() {
	var width, height = termbox.Size()
	for x := 0; x < width; x++ {
		termbox.SetCell(x, height-2, ' ', termbox.ColorBlack, termbox.ColorGreen)
	}
}

func drawMessages(messageTexts *list.List) {
	var width, height = termbox.Size()
	var y = height - 3
	for msg := messageTexts.Back(); msg != nil; msg = msg.Prev() {
		msgStr := msg.Value.(string)
		var linesRequired = len(msgStr)/width + 1

		for ; y >= 0 && linesRequired > 0; y-- {
			var x int
			for charIndex := width * (linesRequired - 1); charIndex < len(msgStr)-1 && x < width; charIndex++ {
				termbox.SetCell(x, y, []rune(msgStr)[charIndex], termbox.ColorWhite, termbox.ColorBlack)
				x++
			}

			linesRequired--
		}
	}
}

func drawIrc(messageTexts *list.List) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawMessages(messageTexts)
	drawSeparator()
	termbox.Flush()
}

func drawInputString(inputString string) {
	var width, height = termbox.Size()
	var linesRequired = len(inputString)/width + 1

	var charIndex int
	for y := height - (1 + linesRequired); y >= 0 && y < height-1; y++ {
		for x := 0; x < width; x++ {
			var c rune
			if charIndex < len(inputString) {
				c = []rune(inputString)[charIndex]
			} else {
				c = ' '
			}
			termbox.SetCell(x, y, c, termbox.ColorBlack, termbox.ColorGreen)

			charIndex++
		}
	}
	termbox.Flush()
}
