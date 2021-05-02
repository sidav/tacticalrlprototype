package main

type sacode uint8

const (
	SA_SLASH = iota
	SA_STRONG_HIT
	SA_MOVEBY
	SA_LUNGE
)

type specialAttack struct {
	code                  sacode
	relativeTargetsCoords [][2]int
}

type specialAttackData struct {
	damage int
}
