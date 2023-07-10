package enemy

import "root/model"

const (
	GOBLIN_CLASS string = "Goblin"
)

type Goblin struct {
	model.Character
}

func GoblinNew() *Goblin {
	return &Goblin{
		Character: model.NewCharacter(model.SetHealth(10, 20), model.SetRandomName, model.SetClass(GOBLIN_CLASS)),
	}
}
