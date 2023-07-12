package weapon

import (
	"fmt"
)

const (
	MINI_GUN_WPN     string = "MiniGun"
	XM1014_WPN       string = "XM1014"
	AK_47_WPN        string = "AK-47"
	M4A4_WPN         string = "M4A4"
	AWP_WPN          string = "AWP"
	DESERT_EAGLE_WPN string = "Desert Eagle"
	P90_WPN          string = "P90"
)

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

func (w Weapon) String() {
	fmt.Printf("Name: %s\nPrecision: %d%%\nLeast Damage: %dHP\nBiggest Damage: %dHP\nAmmunition: %d\n\n",
		w.name,
		w.precision,
		w.leastDamage,
		w.biggestDamage,
		w.ammunition)
}
