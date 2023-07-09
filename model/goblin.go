package model

type Goblin struct {
	Character
}

func GoblinNew() *Goblin {
	return &Goblin{
		Character: NewCharacter(setHealth(10, 20), setRandomName, setClass(GOBLIN_CLASS)),
	}
}
