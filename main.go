package main

import (
	"time"

	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func playSound() {
	file, _ := f.Open("jingle-bells-8bit.mp3")
	streamer, format, _ := mp3.Decode(file)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	vol := &effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   -3,
	}

	speaker.Play(vol)
}

func main() {
	playSound()
	clearConsole()

	printTree(true)

	for {
		resetCursor()
		printTree(false)
		time.Sleep(200 * time.Millisecond)
	}
}
