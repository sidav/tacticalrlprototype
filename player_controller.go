package main

import (
	console "github.com/sidav/golibrl/console"
)

type playerController struct {
}

func (p *playerController) playerControl(d *gameMap) {
	player := d.player
	valid_key_pressed := false
	movex := 0
	movey := 0
	// px, py := CURRENT_MAP.player.getCoords()
	for !valid_key_pressed {
		key_pressed := console.ReadKey()
		valid_key_pressed = true
		movex, movey = p.keyToDirection(key_pressed)
		if movex == 0 && movey == 0 {
			switch key_pressed {
			case "ESCAPE":
				GAME_IS_RUNNING = false
			case "f":
				player.plannedSpecialAttack = &specialAttack{
					code:                  0,
					relativeTargetsCoords: [][2]int{{-1, 0}},
				}
				player.spendTurnsForAction(10)
			default:
				valid_key_pressed = false
				log.AppendMessagef("Unknown key %s (Wrong keyboard layout?)", key_pressed)
			}
		}
	}
	// move player's pawn here and something
	if movex != 0 || movey != 0 {
		CURRENT_MAP.movePawnOrOpenDoorByVector(CURRENT_MAP.player, true, movex, movey)
		player.spendTurnsForAction(10)
	}
}

func (p *playerController) keyToDirection(keyPressed string) (int, int) {
	switch keyPressed {
	case "s", "2":
		return 0, 1
	case "w", "8":
		return 0, -1
	case "a", "4":
		return -1, 0
	case "d", "6":
		return 1, 0
	case "7":
		return -1, -1
	case "9":
		return 1, -1
	case "1":
		return -1, 1
	case "3":
		return 1, 1
	default:
		return 0, 0
	}
}
