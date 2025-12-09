package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func printStars(numberOfStars int, useDelay bool) {
	for range numberOfStars {
		symbol := "*"
		printColor(Bold)

		if numberOfStars == 1 {
			printColor(BrightYellow)

		} else if randomNumber := rand.Intn(100); randomNumber > (100 - сolorBallRate) {
			printRandomColor()
		}

		fmt.Print(symbol)
		fmt.Print(Green)

		if useDelay {
			time.Sleep(starSpawnDelay)
		}
	}
	print("\n")
}

func printBase() {
	for i := 0; i <= lenTree/10; i += 2 {
		pot := strings.Builder{}
		pot.WriteString("▜")
		pot.WriteString(strings.Repeat("█", (lenTree/3)+((lenTree+1)%2)-i))
		pot.WriteString("▛")
		printSpaces((lenTree / 2 * 2) - (pot.Len() / 6))
		printColor(Brown)
		pot.WriteByte('\n')
		print(pot.String())
	}
}
func print(s ...string) {
	for _, el := range s {
		os.Stdout.WriteString(el)
	}
}

func printTree(withDelay bool) {
	print(strings.Repeat("\n", lenTree/5))
	indent := lenTree / 2 * 2
	for n := lenTree; n > 0; n -= 2 {
		printSpaces(indent)
		printStars(lenTree-n+1, withDelay)
		indent--
	}

	printBase()

	print(strings.Repeat("\n", lenTree/5+1), Reset)
}

func printSpaces(numberOfSpaces int) {
	print(strings.Repeat(" ", numberOfSpaces))
}
