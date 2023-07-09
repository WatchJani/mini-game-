package model

import (
	"math/rand"
	"time"
)

type Main struct {
	Character
	PowerAbility
	attack     map[string][]int
	items      map[interface{}]int
	visibility int
}

type PowerAbility struct {
	name         string
	minDamage    int
	maxDamage    int
	coolDown     time.Duration
	numberOfFire int
}

type Item struct {
	name string
}

func NewPowerAbility(name string, min, max int, coolDown time.Duration, numberOfFire int) PowerAbility {
	return PowerAbility{
		name:         name,
		minDamage:    min,
		maxDamage:    max,
		coolDown:     coolDown,
		numberOfFire: numberOfFire,
	}
}

func Empty() PowerAbility {
	return PowerAbility{}
}

func DefaultMain() *Main {
	attack := map[string][]int{
		//preciznost, najmanje moguce napravljene stete, najvise moguce napravljene stete
		NOVA:         {80, 26, 64},
		XM1014:       {80, 20, 30},
		AK_47:        {50, 36, 47},
		M4A4:         {59, 33, 39},
		AWP:          {18, 115, 459},
		DESERT_EAGLE: {20, 34, 46},
		P90:          {23, 21, 26},
	}

	return &Main{
		Character:    NewCharacter(),
		PowerAbility: Empty(),
		attack:       attack,
		items:        make(map[interface{}]int),
		visibility:   rand.Intn(68-12+1) + 12,
	}
}

func NewItem(name string) Item {
	return Item{
		name: name,
	}
}

func (m *Main) NewPowerAbility(ability PowerAbility) {
	m.PowerAbility = ability
}

func (m *Main) AddNewItem(item interface{}) {
	m.items[item]++
}

func (m Main) UseWeapon(weapon string) []int {
	return m.attack[weapon]
}
