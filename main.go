package main

import console "github.com/sidav/golibrl/console"

func main() {
	console.Init_console("", console.TCellRenderer)
	defer console.Close_console()
	game := game{}
	game.runGame()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func euclideanDistance(fx, fy, tx, ty int) int {
	return abs(fx-tx)+abs(fy-ty)
}

