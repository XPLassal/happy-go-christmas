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
	spawnCount := (lenRow * snowDensity) / 100
	newSnowflakes := make([]snowflake, 0, len(snowflakes)+spawnCount)

	for _, v := range snowflakes {
		if v.p.y < lenSide {
			newSnowflakes = append(newSnowflakes, snowflake{Point{v.p.x + (-2 + (rand.Intn(4) + 1)), v.p.y + 1}, v.texture})
		}
	}
	for range spawnCount {
		newSnowflakes = append(newSnowflakes, generateNewSnowflake())
	}
	snowflakes = newSnowflakes
}
