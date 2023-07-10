package main

import (
	"fmt"
	"math/rand"
	c "root/model/hero/class"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	tank := c.NewTank()

	fmt.Println(tank)

	//===================================================================

	// steve := c.DefaultMain()
	// goblin := c.GoblinNew()

	// shadowVeil := c.NewPowerAbility("Shadow Veil", 25, 50, 5, 3)
	// fmt.Println(shadowVeil)

	// fmt.Println(steve.UseWeapon(c.AK_47))
	// fmt.Println(steve)

	// steve.NewPowerAbility(shadowVeil)

	// fmt.Println(steve)

	// apple := c.NewItem("apple")

	// steve.AddNewItem(apple)

	// fmt.Println(steve)

	// spawn := map[string]c.HealthReader{
	// 	"goblin": c.GoblinNew(),
	// }

	// c.GetBothHealth(steve, spawn["goblin"])
}
