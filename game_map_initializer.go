package main

var testMap = []string{
	"########################",
	"#......#...............#",
	"#......#...............#",
	"#......#...............#",
	"#......+...............#",
	"#......#...............#",
	"#......#.......####+####",
	"###+####.......#.......#",
	"#..............#.......#",
	"#..............+.......#",
	"#..............#.......#",
	"#..............#.......#",
	"########################",
}

func (dung *gameMap) initialize_level() { //crap of course
	dung.pawns = make([]*pawn, 0)
	dung.MakeMapFromGenerated(&testMap)
	dung.spawnPlayerAtRandomPosition()
	dung.init_placeItemsAndEnemies()
}

func (dung *gameMap) initTilesArrayForSize(sx, sy int) {
	dung.tiles = make([][]d_tile, sx)
	for i := range dung.tiles {
		dung.tiles[i] = make([]d_tile, sy)
	}
}

func (dung *gameMap) init_placeItemsAndEnemies() {
	dung.spawnPawnAtRandomPosition(PAWN_SWORDSMAN, 1)
	dung.spawnPawnAtRandomPosition(PAWN_SWORDSMAN, 5)
	dung.spawnPawnAtRandomPosition(PAWN_WEAKLING, 15)
}

func (dung *gameMap) MakeMapFromGenerated(generated_map *[]string) {
	levelsizex := len(*generated_map)
	levelsizey := len((*generated_map)[0])
	dung.initTilesArrayForSize(levelsizex, levelsizey)

	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			currDungCell := &dung.tiles[x][y]
			currGenCell := (*generated_map)[x][y] //GetCell(x, y)
			switch currGenCell {
			case '#':
				currDungCell.code = TILE_WALL
			case '.':
				currDungCell.code = TILE_FLOOR
			case '+':
				currDungCell.code = TILE_DOOR
			default:
				currDungCell.code = TILE_UNDEFINED
			}
		}
	}
}

func (dung *gameMap) spawnPlayerAtRandomPosition() {
	CURRENT_MAP.player = &pawn{
		code:          PAWN_PLAYER,
		hp:            5,
		x:             1,
		y:             1,
		nextTurnToAct: 0,
	}
}

func (dung *gameMap) spawnPawnAtRandomPosition(pcode pawnCode, count int) {
	w, h := dung.getSize()
	ai := ai{}
	x, y := 0, 0
	for i := 0; i < count; i++ {
		for !dung.isTilePassableAndNotOccupied(x, y) {
			x = rnd.Rand(w)
			y = rnd.Rand(h)
		}
		dung.pawns = append(dung.pawns, &pawn{
			ai:                   &ai,
			code:                 pcode,
			hp:                   999,
			x:                    x,
			y:                    y,
			nextTurnToAct:        0,
			plannedSpecialAttack: nil,
		})
	}
}

func (dung *gameMap) spawnItemAtRandomPosition(name string, count int) {
}
