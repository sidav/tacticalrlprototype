package main

import "github.com/sidav/golibrl/astar"

func (p *pawn) actAsPawn() {
	if p.ai == nil {
		log.AppendMessage("no ai error")
	}
	// try to move to player
	px, py := p.getCoords()
	tx, ty := CURRENT_MAP.player.getCoords()
	pfmap := CURRENT_MAP.createCostMapForPathfinding()
	path := astar.FindPath(pfmap, px, py, tx, ty, true, 25, true, true)
	movex, movey := path.GetNextStepVector()
	CURRENT_MAP.movePawnOrOpenDoorByVector(p, true, movex, movey)
}

type ai struct {}
