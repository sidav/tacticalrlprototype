package main

type (
	pawn struct {
		code                    pawnCode
		ai                      *ai
		hp, x, y, nextTurnToAct int
		plannedSpecialAttack    *specialAttack
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
		p.plannedSpecialAttack = nil
	}
}
