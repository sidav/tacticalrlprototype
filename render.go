package main

import (
	cw "github.com/sidav/golibrl/console"
)

const (
	FogOfWarColor = cw.DARK_GRAY
)

type renderer struct {
	R_VIEWPORT_WIDTH     int
	R_VIEWPORT_HEIGHT    int
	R_VIEWPORT_CURR_X    int
	R_VIEWPORT_CURR_Y    int
	RENDER_DISABLE_LOS   bool
	R_ALL_INVERSE_COORDS [][2]int
}

func initRenderer() *renderer {
	return &renderer{
		R_VIEWPORT_WIDTH: 40,
		R_VIEWPORT_HEIGHT: 20,
		R_VIEWPORT_CURR_X: 0,
		R_VIEWPORT_CURR_Y: 0,	
	}
}

func (r *renderer) updateBoundsIfNeccessary(force bool) {
	if cw.WasResized() || force {
		cw, ch := cw.GetConsoleSize()
		r.R_VIEWPORT_WIDTH = 2 * cw / 3
		r.R_VIEWPORT_HEIGHT = ch - len(log.Last_msgs) - 1
		//SIDEBAR_X = VIEWPORT_W + 1
		//SIDEBAR_W = cw - VIEWPORT_W - 1
		//SIDEBAR_H = ch - LOG_HEIGHT
		//SIDEBAR_FLOOR_2 = 7  // y-coord right below resources info
		//SIDEBAR_FLOOR_3 = 11 // y-coord right below "floor 2"
	}
}

//func (r *renderer) r_areRealCoordsInViewport(x, y int) bool {
//	return x - r.R_VIEWPORT_CURR_X < r.R_VIEWPORT_WIDTH && y - r.R_VIEWPORT_CURR_Y < r.R_VIEWPORT_HEIGHT
//}

func (r *renderer) r_CoordsToViewport(x, y int) (int, int) {
	vpx, vpy := x-r.R_VIEWPORT_CURR_X, y-r.R_VIEWPORT_CURR_Y
	if vpx >= r.R_VIEWPORT_WIDTH || vpy >= r.R_VIEWPORT_HEIGHT {
		return -1, -1
	}
	return vpx, vpy
}

func (r *renderer) areCoordsInInverseList(x, y int) bool {
	for i := range r.R_ALL_INVERSE_COORDS {
		if r.R_ALL_INVERSE_COORDS[i][0] == x && r.R_ALL_INVERSE_COORDS[i][1] == y {
			return true
		}
	}
	return false
}

func (r *renderer) updateViewportCoords(p *pawn) {
	r.R_VIEWPORT_CURR_X = p.x - r.R_VIEWPORT_WIDTH/2
	r.R_VIEWPORT_CURR_Y = p.y - r.R_VIEWPORT_HEIGHT/2
}

func (r *renderer) refreshInverseCoords() {
	r.R_ALL_INVERSE_COORDS = [][2]int{}
	for _, p := range CURRENT_MAP.pawns {
		if p.plannedSpecialAttack != nil {
			for _, rc := range p.plannedSpecialAttack.relativeTargetsCoords {
				r.R_ALL_INVERSE_COORDS = append(r.R_ALL_INVERSE_COORDS, [2]int{p.x + rc[0], p.y + rc[1]})
			}
		}
	}
	p := CURRENT_MAP.player
	if p.plannedSpecialAttack != nil {
		for _, rc := range p.plannedSpecialAttack.relativeTargetsCoords {
			r.R_ALL_INVERSE_COORDS = append(r.R_ALL_INVERSE_COORDS, [2]int{p.x + rc[0], p.y + rc[1]})
		}
	}
}

func (r *renderer) renderLevel(d *gameMap, flush bool) {
	r.updateBoundsIfNeccessary(false)
	cw.Clear_console()
	r.refreshInverseCoords()
	vismap := *(d.getFieldOfVisionFor(d.player))
	r.updateViewportCoords(d.player)
	// render level. vpx, vpy are viewport coords, whereas x, y are real coords.
	for x := r.R_VIEWPORT_CURR_X; x < r.R_VIEWPORT_CURR_X+r.R_VIEWPORT_WIDTH; x++ {
		for y := 0; y < r.R_VIEWPORT_CURR_Y+r.R_VIEWPORT_HEIGHT; y++ {
			vpx, vpy := r.r_CoordsToViewport(x, y)
			if !d.areCoordinatesValid(x, y) {
				continue
			}
			cellRune := d.tiles[x][y].getAppearance().appearance
			cellColor := d.tiles[x][y].getAppearance().color
			if d.tiles[x][y].isBloody {
				cellColor = cw.DARK_RED
			}
			if r.RENDER_DISABLE_LOS || vismap[x][y] {
				d.tiles[x][y].wasSeenByPlayer = true
				invert := d.tiles[x][y].getAppearance().inverse
				if r.areCoordsInInverseList(x, y) {
					invert = !invert
				}
				if invert {
					cw.SetFgColor(cw.BLACK)
					cw.SetBgColor(cellColor)
				} else {
					cw.SetFgColor(cellColor)
					cw.SetBgColor(cw.BLACK)
				}
			} else {
				if d.tiles[x][y].wasSeenByPlayer {
					if d.tiles[x][y].getAppearance().inverse {
						cw.SetFgColor(cw.BLACK)
						cw.SetBgColor(FogOfWarColor)
					} else {
						cw.SetFgColor(FogOfWarColor)
						cw.SetBgColor(cw.BLACK)
					}
				}
			}
			if d.tiles[x][y].wasSeenByPlayer {
				cw.PutChar(cellRune, vpx, vpy)
			}
			cw.SetBgColor(cw.BLACK)
		}
	}
	//render items
	//for _, item := range d.items {
	//	if r.RENDER_DISABLE_LOS || vismap[item.x][item.y] {
	//		renderItem(item)
	//	}
	//}

	//render pawns
	for _, pawn := range d.pawns {
		if r.RENDER_DISABLE_LOS || vismap[pawn.x][pawn.y] {
			r.renderPawn(pawn, r.areCoordsInInverseList(pawn.x, pawn.y))
		}
	}

	//render projectiles
	//for _, proj := range d.projectiles {
	//	if areCoordinatesValid(proj.x, proj.y) && (r.RENDER_DISABLE_LOS || vismap[proj.x][proj.y]) {
	//		renderProjectile(proj)
	//	}
	//}

	//render player
	r.renderPawn(d.player, false)

	// renderPlayerStats(d)
	r.renderLog(false)
	r.renderUI()

	if flush {
		cw.Flush_console()
	}
}

//func (r *renderer) renderProjectile(p *projectile) {
//	SetColor(RED, BLACK)
//	x, y := r_CoordsToViewport(p.x, p.y)
//	PutChar('*', x, y)
//}

func (r *renderer) renderPawn(p *pawn, inverse bool) {
	app := p.getStaticData().ccell.appearance
	clr := p.getStaticData().ccell.color
	if inverse {
		cw.SetColor(cw.BLACK, clr)
	} else {
		cw.SetFgColor(clr)
	}
	x, y := r.r_CoordsToViewport(p.x, p.y)
	cw.PutChar(app, x, y)
	cw.SetBgColor(cw.BLACK)
}

//func (r *renderer) renderItem(i *i_item) {
//	SetFgColor(i.ccell.color)
//	x, y := r_CoordsToViewport(i.x, i.y)
//	PutChar(i.ccell.appearance, x, y)
//}

//func (r *renderer) renderBullets(currCoords []*routines.Vector, currDirs []*routines.Vector, d *gameMap) {
//	renderLevel(d, false)
//
//	for i:=0; i<len(currCoords); i++ {
//		currx, curry := currCoords[i].GetRoundedCoords()
//		tox, toy := currDirs[i].GetRoundedCoords()
//		SetFgColor(YELLOW)
//		bulletRune := '*'
//		if !d.isPawnPresent(currx, curry) && !d.isTileOpaque(currx, curry) {
//			bulletRune = getTargetingChar(tox, toy)
//		}
//		x, y := r_CoordsToViewport(currx, curry)
//		PutChar(bulletRune, x, y)
//	}
//	Flush_console()
//	time.Sleep(35 * time.Millisecond)
//}

//
// UI-related stuff below
//

//func (r *renderer) renderPlayerStats(d *gameMap) {
//	player := d.player
//	pinv := player.inventory
//	statusbarsWidth := 80 - r.R_VIEWPORT_WIDTH - 3
//
//	hpPercent := player.getHpPercent()
//	var hpColor int
//	switch {
//	case hpPercent < 33:
//		hpColor = RED
//		break
//	case hpPercent < 66:
//		hpColor = YELLOW
//		break
//	default:
//		hpColor = DARK_GREEN
//		break
//	}
//	SetFgColor(hpColor)
//
//	renderStatusbar(fmt.Sprintf("HP: (%d/%d)", player.hp, player.maxhp), player.hp, player.maxhp,
//		r.R_VIEWPORT_WIDTH+1, 0, statusbarsWidth, hpColor)
//
//	if player.wearedArmor == nil {
//		SetFgColor(BEIGE)
//		PutString("No armor", r.R_VIEWPORT_WIDTH+1, 1)
//	} else {
//		SetFgColor(player.wearedArmor.ccell.color)
//		renderStatusbar(fmt.Sprintf("ARMOR: (%d/%d)", player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor),
//			player.wearedArmor.armorData.currArmor, player.wearedArmor.armorData.maxArmor, r.R_VIEWPORT_WIDTH+1, 1, statusbarsWidth, player.wearedArmor.ccell.color)
//	}
//
//	SetFgColor(BEIGE)
//	if player.weaponInHands != nil {
//		renderStatusbar(fmt.Sprintf("%s (%d/%d)", player.weaponInHands.name, player.weaponInHands.weaponData.ammo,
//			player.weaponInHands.weaponData.maxammo), player.weaponInHands.weaponData.ammo,
//			player.weaponInHands.weaponData.maxammo, r.R_VIEWPORT_WIDTH+1, 2, statusbarsWidth, DARK_YELLOW)
//	} else {
//		PutString("Barehanded", r.R_VIEWPORT_WIDTH+1, 2)
//	}
//
//	SetFgColor(BEIGE)
//	PutString(fmt.Sprintf("INV: %d/%d", len(pinv.items), pinv.maxItems), r.R_VIEWPORT_WIDTH+1, 3)
//
//	SetColor(BEIGE, BLACK)
//	ammoLine := fmt.Sprintf("BULL:%d/%d", pinv.ammo[AMMO_BULL], pinv.maxammo[AMMO_BULL])
//	PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 4)
//	ammoLine = fmt.Sprintf("SHLL:%d/%d", pinv.ammo[AMMO_SHEL], pinv.maxammo[AMMO_SHEL])
//	PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 5)
//	ammoLine = fmt.Sprintf("RCKT:%d/%d", pinv.ammo[AMMO_RCKT], pinv.maxammo[AMMO_RCKT])
//	PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 6)
//	ammoLine = fmt.Sprintf("CELL:%d/%d", pinv.ammo[AMMO_CELL], pinv.maxammo[AMMO_CELL])
//	PutString(ammoLine, r.R_VIEWPORT_WIDTH+1, 7)
//
//	timeline := fmt.Sprintf("TIME: %d.%d (%d.%d)", CURRENT_TURN/10, CURRENT_TURN%10,
//		player.playerData.lastSpentTimeAmount/10, player.playerData.lastSpentTimeAmount%10)
//	PutString(timeline, r.R_VIEWPORT_WIDTH+1, 9)
//
//	remEnemiesLine := fmt.Sprintf("ENEMIES LEFT: %d", len(d.pawns))
//	PutString(remEnemiesLine, r.R_VIEWPORT_WIDTH+1, 10)
//}
//
//func (r *renderer) renderTargetingLine(fromx, fromy, tox, toy int, flush bool, d *gameMap) {
//	renderLevel(d, false)
//	line := routines.GetLine(fromx, fromy, tox, toy)
//	char := '?'
//	if len(line) > 1 {
//		dirVector := routines.CreateVectorByStartAndEndInt(fromx, fromy, tox, toy)
//		dirVector.TransformIntoUnitVector()
//		dirx, diry := dirVector.GetRoundedCoords()
//		char = getTargetingChar(dirx, diry)
//	}
//	if fromx == tox && fromy == toy {
//		renderPawn(d.player, true)
//	}
//	for i := 1; i < len(line); i++ {
//		x, y := line[i].X, line[i].Y
//		if d.isPawnPresent(x, y) {
//			renderPawn(d.getPawnAt(x, y), true)
//		} else {
//			SetFgColor(YELLOW)
//			if i == len(line)-1 {
//				char = 'X'
//			}
//			viewx, viewy := r_CoordsToViewport(line[i].X, line[i].Y)
//			PutChar(char, viewx, viewy)
//		}
//	}
//	if flush {
//		Flush_console()
//	}
//}
//
//func (r *renderer) renderStatusbar(name string, curvalue, maxvalue, x, y, width, barColor int) {
//	barTitle := name
//	PutString(barTitle, x, y)
//	barWidth := width - len(name)
//	filledCells := barWidth * curvalue / maxvalue
//	barStartX := x + len(barTitle) + 1
//	for i := 0; i < barWidth; i++ {
//		if i < filledCells {
//			SetFgColor(barColor)
//			PutChar('=', i+barStartX, y)
//		} else {
//			SetFgColor(DARK_BLUE)
//			PutChar('-', i+barStartX, y)
//		}
//	}
//}
//
//func (r *renderer) getTargetingChar(x, y int) rune {
//	if abs(x) > 1 {
//		x /= abs(x)
//	}
//	if abs(y) > 1 {
//		y /= abs(y)
//	}
//	if x == 0 {
//		return '|'
//	}
//	if y == 0 {
//		return '-'
//	}
//	if x*y == 1 {
//		return '\\'
//	}
//	if x*y == -1 {
//		return '/'
//	}
//	return '?'
//}
//
//func (r *renderer) abs(i int) int {
//	if i < 0 {
//		return -i
//	}
//	return i
//}

func (r *renderer) renderLog(flush bool) {
	cw.SetFgColor(cw.RED)
	for i := 0; i < len(log.Last_msgs); i++ {
		cw.PutString(log.Last_msgs[i].Message, 0, r.R_VIEWPORT_HEIGHT+i)
	}
	if flush {
		cw.Flush_console()
	}
}

//func (r *renderer) renderLine(char rune, fromx, fromy, tox, toy int, flush, exceptFirstAndLast bool) {
//	line := routines.GetLine(fromx, fromy, tox, toy)
//	SetFgColor(RED)
//	if exceptFirstAndLast {
//		for i := 1; i < len(line)-1; i++ {
//			PutChar(char, line[i].X, line[i].Y)
//		}
//	} else {
//		for i := 0; i < len(line); i++ {
//			PutChar(char, line[i].X, line[i].Y)
//		}
//	}
//	if flush {
//		Flush_console()
//	}
//}
