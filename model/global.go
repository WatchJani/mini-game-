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
	DEFAULT_SPEED       int    = 100

	MINI_GUN_WPN     string = "MiniGun"
	XM1014_WPN       string = "XM1014"
	AK_47_WPN        string = "AK-47"
	M4A4_WPN         string = "M4A4"
	AWP_WPN          string = "AWP"
	DESERT_EAGLE_WPN string = "Desert Eagle"
	P90_WPN          string = "P90"
)
