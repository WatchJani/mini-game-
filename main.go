package main

import (
	"fmt"
	"math/rand"
	c "root/model"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	steve := c.DefaultMain()
	goblin := c.GoblinNew()

	shadowVeil := c.NewPowerAbility("Shadow Veil", 25, 50, 5, 3)
	fmt.Println(shadowVeil)

	fmt.Println(steve.UseWeapon(c.AK_47))
	fmt.Println(steve)

	steve.NewPowerAbility(shadowVeil)

	fmt.Println(steve)

	apple := c.NewItem("apple")

	steve.AddNewItem(apple)

	fmt.Println(steve)

	c.GetBothHealth(steve, goblin)
}
