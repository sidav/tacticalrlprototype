package main


func (p *pawn) checkAndPerformSpecialAttack() {
	if p.plannedSpecialAttack != nil {
		log.AppendMessage("Wow!")
		// p.plannedSpecialAttack = nil
	}
}

func (p *pawn) setupPassbyAttack() {
	if p.passbyAttack.getRemainingCD() > 10 {
		return
	}
	p.passbyAttack.pawnTargets = []*pawn{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			pAt := CURRENT_MAP.getPawnAt(p.x+x, p.y+y)
			if pAt != nil && pAt != p {
				p.passbyAttack.pawnTargets = append(p.passbyAttack.pawnTargets, pAt)
			}
		}
	}
}

func (p *pawn) performPassbyAttacks() {
	if p.passbyAttack.getRemainingCD() > 0 {
		return
	}
	for _, targ := range p.passbyAttack.pawnTargets {
		if areCoordsNeighbouring(p.x, p.y, targ.x, targ.y, true) {
			log.AppendMessagef("%s passby-attacked %s for 5 damage", p.getStaticData().name, targ.getStaticData().name)
			targ.hp -= 5
			p.passbyAttack.turnToBeAvailableAgain = CURRENT_TURN + p.passbyAttack.getStaticData().cooldown
		}
	}
}
