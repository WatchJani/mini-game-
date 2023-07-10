package class

import (
	"root/model"
	"root/model/hero"
	w "root/model/hero/weapon"
)

// mogao sam samo napisati primitivne tipove podataka bez funkcija, ali mi ovako izgleda ljepse i imam default opciju:D
// mozda ljepse rjesenje bi bilo da sam ovo sve zapisao u neki file i imao samo jedan konstruktor

func NewTank() *hero.Hero {
	return hero.NewMainClass(
		hero.SetNewCharacter(
			model.SetClass("Tank"),
			model.SetHealth(100, 300),
			model.SetName("Šipka"),
			model.SetSpeed(12),
		),
		hero.SetVisibility(12, 20),
	)
}

func NewSniper() *hero.Hero {
	return hero.NewMainClass(
		hero.SetNewCharacter(
			model.SetClass("Sniper"),
			model.SetName("Dragan"),
			model.SetSpeed(56),
		),
		hero.SetVisibility(40, 70),
		hero.SetAttack(w.AWP()),
	)
}

func NewScout() *hero.Hero {
	return hero.NewMainClass(
		hero.SetNewCharacter(
			model.SetClass("Scout"),
			model.SetHealth(60, 80),
			model.SetName("Janko"),
		),
		hero.SetVisibility(60, 90),
	)
}
