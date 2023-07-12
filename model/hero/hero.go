package hero

import (
	"math/rand"
	"root/model"
	p "root/model/hero/ability"
	w "root/model/hero/weapon"
)

type MainClass func(*Hero)

type Hero struct {
	model.Character
	p.PowerAbility
	attack     map[string]w.Weapon
	items      map[interface{}]int
	visibility int
}

func DefaultHero() *Hero {
	return &Hero{
		Character:    model.NewCharacter(),
		PowerAbility: p.Empty(),
		attack:       w.MiniGun(),
		items:        make(map[interface{}]int),
		visibility:   rand.Intn(68-12+1) + 12,
	}
}

func NewMainClass(mainFn ...MainClass) *Hero {
	main := DefaultHero()

	for _, fn := range mainFn {
		fn(main)
	}

	return main
}

func SetNewCharacter(fn ...model.CharacterFunc) MainClass {
	return func(m *Hero) {
		m.Character = model.NewCharacter(fn...)
	}
}

func SetPowerAbility(ability p.PowerAbility) MainClass {
	return func(m *Hero) {
		m.PowerAbility = ability
	}
}

func SetAttack(attack map[string]w.Weapon) MainClass {
	return func(m *Hero) {
		m.attack = attack
	}
}

func SetItems(items map[interface{}]int) MainClass {
	return func(m *Hero) {
		m.items = items
	}
}

func SetVisibility(min, max int) MainClass {
	return func(m *Hero) {
		m.visibility = rand.Intn(max-min+1) + min
	}
}

func (m *Hero) NewPowerAbility(ability p.PowerAbility) {
	m.PowerAbility = ability
}

func (m *Hero) AddNewItem(item interface{}) {
	m.items[item]++
}

func (m *Hero) UseWeapon(weapon string) w.Weapon {
	return m.attack[weapon]
}

func (m *Hero) TakeWeapon(newWeapon map[string]w.Weapon) {
	m.attack = newWeapon
}
