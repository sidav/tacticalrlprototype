package main

import . "github.com/sidav/golibrl/console"
import "fmt"

func (r *renderer) renderUI() {
	r.renderPlayerStats(&CURRENT_MAP)
}

func (r *renderer) renderPlayerStats(d *gameMap) {
	player := d.player
	cw, _ := GetConsoleSize()
	statusbarsWidth :=  cw - r.R_VIEWPORT_WIDTH - 3

	hpPercent := player.getHpPercent()
	var hpColor int
	switch {
	case hpPercent < 33:
		hpColor = RED
		break
	case hpPercent < 66:
		hpColor = YELLOW
		break
	default:
		hpColor = DARK_GREEN
		break
	}
	SetFgColor(hpColor)

	r.renderStatusbar(fmt.Sprintf("HP: (%d/%d)", player.hp, player.getStaticData().maxhp),
		player.hp, player.getStaticData().maxhp,
		r.R_VIEWPORT_WIDTH+1, 0, statusbarsWidth, hpColor)

	log.AppendMessagef("%d:%d:%d",
		CURRENT_TURN,
		player.passbyAttack.getRemainingCD(),
		player.passbyAttack.getStaticData().cooldown)
	r.renderStatusbar(fmt.Sprintf("Pass-by attack: (%d/%d)",
						player.passbyAttack.getRemainingCD(),
						player.passbyAttack.getStaticData().cooldown),
		player.passbyAttack.getRemainingCD(),
		player.passbyAttack.getStaticData().cooldown,
		r.R_VIEWPORT_WIDTH+1, 1, statusbarsWidth, BLUE)

	//
	//if player.wearedArmor == nil {
	//	SetFgColor(BEIGE)
	//	PutString("No armor", r.R_VIEWPORT_WIDTH+1, 1)
	//} else {
	//	SetFgColor(player.wearedArmor.ccell.color)
	//	renderStatusbar(fmt.Sprintf("ARMOR: (%d/%d)", player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor),
	//		player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor, r.R_VIEWPORT_WIDTH+1, 1, statusbarsWidth, player.wearedArmor.ccell.color)
	//}
	//
	//SetFgColor(BEIGE)
	//if player.weaponInHands != nil {
	//	renderStatusbar(fmt.Sprintf("%s (%d/%d)", player.weaponInHands.name, player.weaponInHands.weaponData.ammo,
	//		player.weaponInHands.weaponData.maxammo), player.weaponInHands.weaponData.ammo,
	//		player.weaponInHands.weaponData.maxammo, r.R_VIEWPORT_WIDTH+1, 2, statusbarsWidth, DARK_YELLOW)
	//} else {
	//	PutString("Barehanded", r.R_VIEWPORT_WIDTH+1, 2)
	//}
	//
	//SetFgColor(BEIGE)
	//PutString(fmt.Sprintf("INV: %d/%d", len(pinv.items), pinv.maxItems), r.R_VIEWPORT_WIDTH+1, 3)
	//
	//SetColor(BEIGE, BLACK)
	//ammoLine := fmt.Sprintf("BULL:%d/%d", pinv.ammo[AMMO_BULL], pinv.maxammo[AMMO_BULL])
	//PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 4)
	//ammoLine = fmt.Sprintf("SHLL:%d/%d", pinv.ammo[AMMO_SHEL], pinv.maxammo[AMMO_SHEL])
	//PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 5)
	//ammoLine = fmt.Sprintf("RCKT:%d/%d", pinv.ammo[AMMO_RCKT], pinv.maxammo[AMMO_RCKT])
	//PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 6)
	//ammoLine = fmt.Sprintf("CELL:%d/%d", pinv.ammo[AMMO_CELL], pinv.maxammo[AMMO_CELL])
	//PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 7)
	//
	//timeline := fmt.Sprintf("TIME: %d.%d (%d.%d)", CURRENT_TURN/10, CURRENT_TURN%10,
	//	player.playerData.lastSpentTimeAmount/10, player.playerData.lastSpentTimeAmount%10)
	//PutString(timeline, r.R_VIEWPORT_WIDTH+1, 9)
	//
	//remEnemiesLine := fmt.Sprintf("ENEMIES LEFT: %d", len(d.pawns))
	//PutString(remEnemiesLine, r.R_VIEWPORT_WIDTH+1, 10)
}

func (r *renderer) renderStatusbar(name string, curvalue, maxvalue, x, y, width, barColor int) {
	barTitle := name
	PutString(barTitle, x, y)
	barWidth := width - len(name)
	filledCells := barWidth * curvalue / maxvalue
	barStartX := x + len(barTitle) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			SetFgColor(barColor)
			PutChar('=', i+barStartX, y)
		} else {
			SetFgColor(DARK_BLUE)
			PutChar('-', i+barStartX, y)
		}
	}
}
