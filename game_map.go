package main

type gameMap struct {
	player      *pawn
	tiles       [][]d_tile
	pawns       []*pawn
	//items       []*i_item
	//projectiles []*projectile
}

func (dung *gameMap) getSize() (int, int) {
	return len(dung.tiles), len(dung.tiles[0])
}

func (dung *gameMap) areCoordinatesValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(dung.tiles) && y < len(dung.tiles[0])
}

func (dung *gameMap) isPawnPresent(ix, iy int) bool {
	x, y := dung.player.x, dung.player.y
	if ix == x && iy == y {
		return true
	}
	for i := 0; i < len(dung.pawns); i++ {
		x, y = dung.pawns[i].x, dung.pawns[i].y
		if ix == x && iy == y {
			return true
		}
	}
	return false
}

func (dung *gameMap) getPawnAt(x, y int) *pawn {
	px, py := dung.player.x, dung.player.y
	if px == x && py == y {
		return dung.player
	}
	for i := 0; i < len(dung.pawns); i++ {
		px, py = dung.pawns[i].x, dung.pawns[i].y
		if px == x && py == y {
			return dung.pawns[i]
		}
	}
	return nil
}

func (dung *gameMap) spitBloodAt(xx, yy int) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if rnd.OneChanceFrom(4) {
				CURRENT_MAP.tiles[xx+x][yy+y].isBloody = true
			}
		}
	}
	CURRENT_MAP.tiles[xx][yy].isBloody = true
}

func (d *gameMap) removePawn(p *pawn) {
	for i := 0; i < len(d.pawns); i++ {
		if p == d.pawns[i] {
			d.pawns = append(d.pawns[:i], d.pawns[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (dung *gameMap) isTilePassable(x, y int) bool {
	if !dung.areCoordinatesValid(x, y) {
		return false
	}
	return dung.tiles[x][y].isPassable()
}

func (dung *gameMap) isTileOpaque(x, y int) bool {
	if !dung.areCoordinatesValid(x, y) {
		return true
	}
	return dung.tiles[x][y].isOpaque()
}

func (dung *gameMap) isTileADoor(x, y int) bool {
	if !dung.areCoordinatesValid(x, y) {
		return false
	}
	return dung.tiles[x][y].isDoor()
}

func (dung *gameMap) openDoor(x, y int) {
	if !dung.areCoordinatesValid(x, y) {
		return
	}
	dung.tiles[x][y].isOpened = true
}

// true if action has been commited
func (dung *gameMap) movePawnOrOpenDoorByVector(p *pawn, mayOpenDoor bool, vx, vy int) bool {
	x, y := p.getCoords()
	x += vx; y += vy
	if dung.isTilePassableAndNotOccupied(x, y) {
		p.x = x; p.y = y
		p.setupPassbyAttack()
		p.spendTurnsForAction(p.getStaticData().moveDelay)
		return true
	}
	p.setupPassbyAttack()
	if dung.isTileADoor(x, y) && mayOpenDoor {
		dung.tiles[x][y].isOpened = true
		p.spendTurnsForAction(p.getStaticData().moveDelay)
		return true
	}
	return false
}

func (dung *gameMap) createCostMapForPathfinding() *[][]int {
	width, height := len(dung.tiles), len((dung.tiles)[0])

	costmap := make([][]int, width)
	for j := range costmap {
		costmap[j] = make([]int, height)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// TODO: optimize by iterating through pawns separately
			if !dung.tiles[i][j].isPassable() || dung.getPawnAt(i, j) != nil {
				costmap[i][j] = -1
			}
		}
	}
	return &costmap
}

func (dung *gameMap) isTilePassableAndNotOccupied(x, y int) bool {
	return dung.isTilePassable(x, y) && !dung.isPawnPresent(x, y)
}