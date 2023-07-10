package model

import (
	"math/rand"
	"time"
)

type MainClass func(*Main)

type Main struct {
	Character
	PowerAbility
	attack     map[string]Weapon
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

type Weapon struct {
	name          string
	precision     int
	leastDamage   int
	biggestDamage int
	ammunition    int
	slot          bool
}

func NewWeapon(name string, precision, leastDamage, biggestDamage, ammunition int) map[string]Weapon {

	weapon := make(map[string]Weapon)

	weapon[name] = Weapon{
		name:          name,
		precision:     precision,
		leastDamage:   leastDamage,
		biggestDamage: biggestDamage,
		ammunition:    ammunition,
		slot:          true,
	}

	return weapon
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
	return &Main{
		Character:    NewCharacter(),
		PowerAbility: Empty(),
		attack:       HouLongHeater(),
		items:        make(map[interface{}]int),
		visibility:   rand.Intn(68-12+1) + 12,
	}
}

//=======================================================================

func NewMainClass(mainFn ...MainClass) *Main {
	main := DefaultMain()

	for _, fn := range mainFn {
		fn(main)
	}

	return main
}

func setNewCharacter(fn ...CharacterFunc) MainClass {
	return func(m *Main) {
		m.Character = NewCharacter(fn...)
	}
}

func setPowerAbility(ability PowerAbility) MainClass {
	return func(m *Main) {
		m.PowerAbility = ability
	}
}

func setAttack(attack map[string]Weapon) MainClass {
	return func(m *Main) {
		m.attack = attack
	}
}

func setItems(items map[interface{}]int) MainClass {
	return func(m *Main) {
		m.items = items
	}
}

func setVisibility(min, max int) MainClass {
	return func(m *Main) {
		m.visibility = rand.Intn(max-min+1) + min
	}
}

// ===================================================================

// mogao sam samo napisati primitivne tipove podataka bez funkcija, ali mi ovako izgleda ljepse i imam default opciju:D
// mozda ljepse rjesenje bi bilo da sam ovo sve zapisao u neki file i imao samo jedan konstruktor
func NewTank() *Main {
	return NewMainClass(
		setNewCharacter(
			setClass("Tank"),
			setHealth(100, 300),
			setName("Šipka"),
			setSpeed(12),
		),
		setVisibility(12, 20),
	)
}

func NewSniper() *Main {
	return NewMainClass(
		setNewCharacter(
			setClass("Sniper"),
			setName("Dragan"),
			setSpeed(56),
		),
		setVisibility(40, 70),
		setAttack(AWP()),
	)
}

func NewScout() *Main {
	return NewMainClass(
		setNewCharacter(
			setClass("Scout"),
			setHealth(60, 80),
			setName("Janko"),
		),
		setVisibility(60, 90),
	)
}

// ====================================================================

func AWP() map[string]Weapon {
	return NewWeapon(AWP_WPN, 12, 100, 300, 15)
}

func HouLongHeater() map[string]Weapon {
	return NewWeapon(MINI_GUN_WPN, 20, 2, 7, 100)
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

func (m Main) UseWeapon(weapon string) Weapon {
	return m.attack[weapon]
}
