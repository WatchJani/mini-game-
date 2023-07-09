package model

import "fmt"

type Moves interface {
	Attack()
}

type HealthReader interface {
	GetHealth() int
	GetName() string
	GetClass() string
}

func GetBothHealth(main, enemy HealthReader) {
	fmt.Printf("%s[%s] health: %d <==> %s[%s] health: %d\n",
		main.GetName(),
		main.GetClass(),
		main.GetHealth(),
		enemy.GetName(),
		enemy.GetClass(),
		enemy.GetHealth(),
	)
}

const (
	GOBLIN_CLASS string = "Goblin"
	MAIN_CLASS   string = "Main Character"

	MAIN_CHARACTER_NAME string = "Steve"
	DEFAULT_HEALTH      int    = 100

	NOVA         string = "Nova"
	XM1014       string = "XM1014"
	AK_47        string = "AK-47"
	M4A4         string = "M4A4"
	AWP          string = "AWP"
	DESERT_EAGLE string = "Desert Eagle"
	P90          string = "P90"
)
