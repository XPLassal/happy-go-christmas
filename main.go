package main

import (
	"os"
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

	cfg, exists := LoadConfig()

	if len(os.Args) > 1 {
		if os.Args[1] == "--edit" {
			cfg = CreateConfig()
			exists = true
		}
	}

	if !exists {
		cfg = CreateConfig()
	}

	handleConfig(cfg)

	clearConsole()

	printTree(cfg, true)

	for {
		resetCursor()
		printTree(cfg, false)
		updateSnow()
		time.Sleep(200 * time.Millisecond)
	}
}
