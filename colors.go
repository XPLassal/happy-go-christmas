package main

import (
	"math/rand"
)

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"

	Red           = "\033[31m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	White         = "\033[37m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Magenta       = "\033[35m"
	Cyan          = "\033[36m"
	Brown         = "\033[0;33m"
)

func resetCursor() {
	print("\033[H")
}

func clearConsole() {
	print("\033[H\033[2J")
}

func printColor(color string) {
	print(color)
}

func getRandomColor() string {
	return BrightColors[rand.Intn(len(BrightColors))]
}

func printRandomColor() string {
	randomColor := getRandomColor()
	return randomColor
}
