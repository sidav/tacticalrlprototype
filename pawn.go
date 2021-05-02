package main

type (
	pawn struct {
		code                    pawnCode
		ai                      *ai
		hp, x, y, nextTurnToAct int

		plannedSpecialAttack    *specialAttack
		passbyAttack            *passbyAttack
	}
)

func (p *pawn) isDead() bool {
	return p.hp <= 0
}

func (p *pawn) spendTurnsForAction(turns int) {
	p.nextTurnToAct = CURRENT_TURN + turns
}

func (p *pawn) isTimeToAct() bool {
	return p.nextTurnToAct <= CURRENT_TURN
}

func (p *pawn) getStaticData() *pawnStaticData {
	return pawnsStaticData[p.code]
}

func (p *pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) getHpPercent() int {
	return p.hp * 100 / p.getStaticData().maxhp
}

func (p *pawn) checkAndPerformSpecialAttack() {
	if p.plannedSpecialAttack != nil {
		log.AppendMessage("Wow!")
		// p.plannedSpecialAttack = nil
	}
}

func (p *pawn) setupPassbyAttack() {
	p.passbyAttack.pawnsInRangeAtPrevTurn = []*pawn{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			pAt := CURRENT_MAP.getPawnAt(p.x+x, p.y+y)
			if pAt != nil && pAt != p {
				p.passbyAttack.pawnsInRangeAtPrevTurn = append(p.passbyAttack.pawnsInRangeAtPrevTurn, pAt)
			}
		}
	}
}

func (p *pawn) performPassbyAttacks() {
	for _, targ := range p.passbyAttack.pawnsInRangeAtPrevTurn {
		if euclideanDistance(p.x, p.y, targ.x, targ.y) <= 2 {
			log.AppendMessagef("%s passby-attacked %s for 5 damage", p.getStaticData().name, targ.getStaticData().name)
			targ.hp -= 5 
		}
	}
}
