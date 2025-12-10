package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func printStars(numberOfStars int, useDelay bool, cfg Config) {
	for range numberOfStars {
		symbol := ""

		if cfg.UseLeaf {
			symbol = getRandomEl(leafTexture)
		} else {
			symbol = "*"
		}

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
}

func printBase(p Point) {
	i := 0
	j := barkHight
	for {
		pot := strings.Builder{}

		pot.WriteString("▟")
		pot.WriteString(strings.Repeat("█", ((lenTree-(lenTree+1)%2)/5)+i*2))
		pot.WriteString("▙")

		printSpaces(p, lenRow/2-pot.Len()/6-1)
		printColor(Brown)
		pot.WriteByte('\n')

		print(pot.String())
		i++
		j--

		if j <= 0 {
			break
		}
	}
}
func print(s ...string) {
	for _, el := range s {
		os.Stdout.WriteString(el)
	}
}

func printTree(cfg Config, withDelay bool) {
	y = 0
	for range marginBottom {
		printSpaces(Point{0, y}, lenRow)
		print("\n")
		y++
	}

	i := 1
	for range treeHight {
		printSpaces(Point{0, y}, (marginSide+(lenTree-i))/2)
		printStars(i, withDelay, cfg)
		printSpaces(Point{marginSide + i, y}, marginSide+lenTree)
		print("\n")
		y++
		i += 2
	}

	printBase(Point{0, y})
}

func printSpaces(p Point, endX int) {
	for ; p.x <= endX; p.x++ {
		if texture, ok := isConsistSnowflace(p); ok {
			printSnowflake(texture)
		} else {
			print(" ")
		}
	}
}

func printSnowflake(s string) {
	print(Bold, White, s)
}

func getRandomEl(l []string) string {
	return l[rand.Intn(len(l))]
}
