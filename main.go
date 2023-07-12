package main

import (
	"fmt"
	"math/rand"
	"root/model"
	e "root/model/enemy"
	h "root/model/hero"
	c "root/model/hero/class"
	w "root/model/hero/weapon"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	main := map[string]*h.Hero{
		"tank":   c.NewTank(),
		"sniper": c.NewSniper(),
		"scout":  c.NewScout(),
	}

	hero := main["tank"]

	weapon := map[string]map[string]w.Weapon{
		w.AWP_WPN:      w.AWP(),
		w.MINI_GUN_WPN: w.MiniGun(),
	}

	AWP := weapon[w.AWP_WPN]

	hero.TakeWeapon(AWP)

	AWP[w.AWP_WPN].String()

	fmt.Println(hero)

	enemy := map[string]model.HealthReader{
		"goblin": e.GoblinNew(),
	}

	model.GetBothHealth(hero, enemy["goblin"])

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
