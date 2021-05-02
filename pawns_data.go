package main

import cw "github.com/sidav/golibrl/console"

type pawnCode uint8

const (
	PAWN_WEAKLING = iota
	PAWN_PLAYER
	PAWN_SWORDSMAN
)

type pawnStaticData struct {
	name  string
	maxhp int
	ccell consoleCell
}

var pawnsStaticData = map[pawnCode]*pawnStaticData{
	PAWN_WEAKLING: {
		name:  "Weakling",
		maxhp: 1,
		ccell: consoleCell{
			appearance: 'w',
			color:      cw.BEIGE,
			inverse:    false,
		}},
	PAWN_PLAYER: {
		name:  "Player",
		maxhp: 10,
		ccell: consoleCell{
			appearance: '@',
			color:      cw.BEIGE,
			inverse:    false,
		}},
	PAWN_SWORDSMAN: {
		name:  "Swordsman",
		maxhp: 5,
		ccell: consoleCell{
			appearance: 's',
			color:      cw.BEIGE,
			inverse:    false,
		}},
}
