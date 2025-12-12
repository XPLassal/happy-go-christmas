package main

import (
	"math/rand"
	"os"
	"strings"
)

func getStars(numberOfStars int, cfg Config) string {
	sb := CustomStringsBuilder{}
	sb.Grow(numberOfStars * 25)

	for range numberOfStars {
		symbol := ""

		if cfg.UseLeaf {
			symbol = getRandomEl(leafTexture)
		} else {
			symbol = "*"
		}

		sb.WriteString(Bold)

		if numberOfStars == 1 {
			sb.WriteString(BrightYellow)

		} else if randomNumber := rand.Intn(100); randomNumber > (100 - сolorBallRate) {
			sb.WriteString(getRandomColor())
		}

		sb.WriteStrings(symbol, Green)
	}
	sb.WriteStrings(Reset)
	return sb.String()
}

func getBase(p Point) (int, string) {
	sb := CustomStringsBuilder{}
	sb.Grow(barkHight * lenRow * 10)

	baseWidth := max(lenTree/3, 3)

	if baseWidth%2 == 0 {
		baseWidth--
	}

	i := 0
	j := max(barkHight, 1)

	for {
		currentWidth := baseWidth + (i * 2)

		fillCount := currentWidth - 2
		totalTreeWidth := marginSide + lenTree
		padding := (totalTreeWidth - currentWidth) / 2

		sb.WriteStrings(getSpaces(p, padding))
		sb.WriteStrings(Brown, "▟", strings.Repeat("█", fillCount), "▙", Reset)
		sb.WriteStrings(getSpaces(Point{padding + fillCount, p.y}, lenRow-2), "\n")

		p.y++

		i++
		j--
		if j <= 0 {
			break
		}
	}

	return p.y, sb.String()
}

func print(s ...string) {
	for _, el := range s {
		os.Stdout.WriteString(el)
	}
}

func getTree(cfg Config) string {
	sb := CustomStringsBuilder{}
	estimatedSize := (lenSide + marginBottom*2) * lenRow * 10
	sb.Grow(estimatedSize)

	y = 0
	for range marginBottom {
		sb.WriteStrings(getSpaces(Point{0, y}, lenRow), "\n")
		y++
	}

	i := 1
	for range treeHight {
		sb.WriteStrings(getSpaces(Point{0, y}, (marginSide+(lenTree-i))/2))
		sb.WriteStrings(getStars(i, cfg))
		sb.WriteStrings(getSpaces(Point{(marginSide+(lenTree-i))/2 + i, y}, lenRow), "\n")

		y++
		i += 2
	}

	y, str := getBase(Point{0, y})
	sb.WriteStrings(str)

	for range marginBottom {
		sb.WriteStrings(getSpaces(Point{0, y}, lenRow), "\n")
		y++
	}

	return sb.String()
}

func getSpaces(p Point, endX int) string {
	sb := CustomStringsBuilder{}
	sb.Grow((endX - p.x) * 4)
	for ; p.x <= endX; p.x++ {
		if texture, ok := isConsistSnowflace(p); ok {
			sb.WriteString(texture)
		} else {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func getRandomEl(l []string) string {
	return l[rand.Intn(len(l))]
}
