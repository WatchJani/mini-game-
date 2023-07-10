package ability

import "time"

type PowerAbility struct {
	name         string
	minDamage    int
	maxDamage    int
	coolDown     time.Duration
	numberOfFire int
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


