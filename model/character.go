package model

import "math/rand"

const (
	MAIN_CLASS          string = "Main Character"
	MAIN_CHARACTER_NAME string = "Steve"
	DEFAULT_HEALTH      int    = 100
	DEFAULT_SPEED       int    = 100
)

type CharacterFunc func(*Character)

type Character struct {
	health int
	name   string
	class  string
	speed  int
}

var CHARACTER_NAME []string = []string{
	"Gannon", "Sephiroth", "Bowser", "Ganondorf", "Dracula", "Reaper", "Nemesis", "Diablo", "General RAAM",
	"Shao Kahn", "Ridley", "M. Bison", "Albert Wesker", "Big Daddy", "Pyramid Head", "Vaas Montenegro",
	"Arthas Menethil", "GLaDOS", "Handsome Jack", "Alduin", "The Joker", "Ultron", "King K. Rool",
	"Magus", "Malthael", "Mother Brain", "Akuma", "The Witch King", "Ornstein and Smough", "Gruntilda",
}

func NewCharacter(opts ...CharacterFunc) Character {
	character := defaultCharacter()

	for _, fn := range opts {
		fn(&character)
	}

	return character
}

func defaultCharacter() Character {
	return Character{
		health: DEFAULT_HEALTH,
		name:   MAIN_CHARACTER_NAME,
		class:  MAIN_CLASS,
		speed:  DEFAULT_SPEED,
	}
}

func SetHealth(min, max int) CharacterFunc {
	return func(c *Character) {
		c.health = rand.Intn(max-min+1) + min
	}
}

func SetSpeed(speed int) CharacterFunc {
	return func(c *Character) {
		c.speed = speed
	}
}

func SetName(name string) CharacterFunc {
	return func(c *Character) {
		c.name = name
	}
}

func SetRandomName(c *Character) {
	c.name = CHARACTER_NAME[rand.Intn(len(CHARACTER_NAME))]
}

func SetClass(class string) CharacterFunc {
	return func(c *Character) {
		c.class = class
	}
}

func (c Character) GetClass() string {
	return c.class
}

func (c Character) GetHealth() int {
	return c.health
}

func (c Character) GetName() string {
	return c.name
}
