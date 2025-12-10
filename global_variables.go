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

	leafTexture       = []string{"⣿", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾"}
	snowflakesTexture = []string{"*", ".", "❄", "·", "✽", "❅", "❋"}

	snowflakes = make([]snowflake, 0, lenTree*lenTree)

	y = 0

	//go:embed jingle-bells-8bit.mp3
	f embed.FS

	lenTree,
	marginBottom,
	barkHight,
	treeHight,
	marginSide,
	lenSide,
	lenRow,
	snowDensity int
)

const (
	сolorBallRate  = 35
	starSpawnDelay = 10 * time.Millisecond
)
