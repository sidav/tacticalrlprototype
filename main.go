package main

import console "github.com/sidav/golibrl/console"

func main() {
	console.Init_console("", console.TCellRenderer)
	defer console.Close_console()
	game := game{}
	game.runGame()
}
