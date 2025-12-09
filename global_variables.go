package main

import (
	"embed"
	"time"
)

var (
	BrightColors = []string{
		Red,
		BrightRed,
		BrightYellow,
		BrightBlue,
		White,
		BrightMagenta,
		BrightCyan,
		Yellow,
		Blue,
		Magenta,
		Cyan,
		BrightGreen,
	}

	leaf           = [...]string{"⣿", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾"}
	toys           = [...]string{"❄", "❅", "❆", "✼", "✽"}
	starSpawnDelay = 10 * time.Millisecond
	//go:embed jingle-bells-8bit.mp3
	f embed.FS
)

const (
	lenTree       = 25
	сolorBallRate = 35
)
