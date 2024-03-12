package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func roll() int {
	return rand.Intn(6) + 1
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	lowest := 6
	sum := 0
	for i := 0; i < 4; i++ {
		roll := roll()
		sum += roll
		if roll < lowest {
			lowest = roll

		}
	}
	return sum - lowest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	hitpoints := 10 + Modifier(char.Constitution)
	char.Hitpoints = hitpoints
	return char
}
