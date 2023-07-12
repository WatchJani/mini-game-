package weapon

func AWP() map[string]Weapon {
	return NewWeapon(AWP_WPN, 12, 100, 300, 15)
}

func MiniGun() map[string]Weapon {
	return NewWeapon(MINI_GUN_WPN, 20, 2, 7, 100)
}
