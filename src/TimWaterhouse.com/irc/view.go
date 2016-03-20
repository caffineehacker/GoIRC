package main

import (
	"container/list"

	"github.com/nsf/termbox-go"
)

//text string

/*func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}*/

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

func drawChar(key termbox.Key) {
	termbox.SetCell(0, 0, 'A', termbox.ColorBlack, termbox.ColorGreen)
	termbox.Flush()
}
