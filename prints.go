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
	baseWidth := max(lenTree/3, 3)

	if baseWidth%2 == 0 {
		baseWidth--
	}

	i := 0
	j := max(barkHight, 1)

	for {
		currentWidth := baseWidth + (i * 2)

		fillCount := currentWidth - 2

		pot := strings.Builder{}
		pot.WriteString("▟")
		pot.WriteString(strings.Repeat("█", fillCount))
		pot.WriteString("▙")
		pot.WriteByte('\n')
		totalTreeWidth := marginSide + lenTree
		padding := (totalTreeWidth - currentWidth) / 2

		printSpaces(p, padding)

		print(Brown, pot.String())

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

	for range marginBottom {
		printSpaces(Point{0, y}, lenRow)
		print("\n")
		y++
	}
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
