package main

type sacode uint8

const (
	SA_SLASH = iota
	SA_STRONG_HIT
	SA_PASSBY
	SA_LUNGE
)

type specialAttack struct {
	code                   sacode
	turnToBeAvailableAgain int
	relativeTargetsCoords  [][2]int
	pawnTargets            []*pawn
}

func (sa *specialAttack) getRemainingCD() int {
	cdr := sa.turnToBeAvailableAgain - CURRENT_TURN
	if cdr < 0 {
		cdr = 0
	}
	return cdr///10
}

func (sa *specialAttack) getStaticData() *specialAttackData {
	return specialAttacksData[sa.code]
}

type specialAttackData struct {
	damage   int
	cooldown int
}

var specialAttacksData = map[sacode]*specialAttackData{
	SA_PASSBY: {
		damage:   5,
		cooldown: 70,
	},
}
