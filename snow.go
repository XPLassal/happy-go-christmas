package main

import (
	"math/rand"
)

type snowflake struct {
	p       Point
	texture string
}

func isConsistSnowflace(p Point) (string, bool) {
	for _, v := range snowflakes {
		if p.isEquale(v.p) {
			return v.texture, true
		}
	}
	return "", false
}

func generateNewSnowflake() snowflake {
	return snowflake{Point{x: rand.Intn(lenRow), y: 0}, getRandomEl(snowflakesTexture)}
}

func updateSnow() {
	newSnowflakes := make([]snowflake, 0, len(snowflakes))
	for _, v := range snowflakes {
		if v.p.y < lenSide {
			newSnowflakes = append(newSnowflakes, snowflake{Point{v.p.x + (-2 + (rand.Intn(4) + 1)), v.p.y + 1}, v.texture})
		}
	}
	snowflakes = append(newSnowflakes, generateNewSnowflake())
	snowflakes = append(newSnowflakes, generateNewSnowflake())
	snowflakes = append(newSnowflakes, generateNewSnowflake())
}
