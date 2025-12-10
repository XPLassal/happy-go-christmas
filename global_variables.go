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

	lenTree      = 0
	marginBottom = lenTree / 5
	barkHight    = lenTree / 10
	treeHight    = (lenTree + 1) / 2
	marginSide   = marginBottom * 3
	lenSide      = treeHight + marginBottom + barkHight
	lenRow       = marginSide + lenTree
)

const (
	сolorBallRate  = 35
	starSpawnDelay = 10 * time.Millisecond
)
