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
