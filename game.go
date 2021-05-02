package main

import (
	"github.com/sidav/golibrl/random/additive_random"
	log2 "gorltemplate/game_log"
)

var (
	GAME_IS_RUNNING bool
	log             log2.GameLog
	rnd             additive_random.FibRandom
	pc              playerController
	CURRENT_TURN    int
	CURRENT_MAP     gameMap
)

type game struct {
}

func areCoordinatesInRangeFrom(fx, fy, tx, ty, srange int) bool {
	return (tx-fx)*(tx-fx)+(ty-fy)*(ty-fy) < srange*srange
}

func (g *game) runGame() {
	log = log2.GameLog{}
	log.Init(5)
	rnd = additive_random.FibRandom{}
	rnd.InitDefault()

	GAME_IS_RUNNING = true
	CURRENT_MAP = gameMap{}
	CURRENT_MAP.initialize_level()
	log.AppendMessage("Init complete")

	rend := initRenderer()

	for GAME_IS_RUNNING {
		rend.renderLevel(&CURRENT_MAP, true)
		CURRENT_MAP.player.checkAndPerformSpecialAttack()
		for CURRENT_MAP.player.isTimeToAct() && GAME_IS_RUNNING {
			pc.playerControl(&CURRENT_MAP)
		}

		// check if pawns should be removed
		for i := 0; i < len(CURRENT_MAP.pawns); i++ {
			if CURRENT_MAP.pawns[i].isTimeToAct() {
				// act for pawns here
				CURRENT_MAP.pawns[i].actAsPawn()
			}
		}
		CURRENT_TURN++
	}
}