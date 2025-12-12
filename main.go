package main

import (
	"os"
	"slices"
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

func haveArg(s string) bool {
	return slices.Contains(os.Args[1:], s)
}

func main() {
	cfg, exists := LoadConfig()
	if haveArg("--edit") {
		cfg = CreateConfig()
		exists = true
	}
	if !exists {
		cfg = CreateConfig()
	}

	handleConfig(cfg)
	clearConsole()

	if haveArg("--static") {
		for range lenTree + marginBottom*2 {
			updateSnow()
		}
		print(getTree(cfg))
		return
	}

	delay := 1 * time.Microsecond
	for _, el := range getTree(cfg) {
		print(string(el))
		time.Sleep(delay)
	}

	for {
		resetCursor()
		print(getTree(cfg))
		updateSnow()
		time.Sleep(200 * time.Millisecond)
	}
}
