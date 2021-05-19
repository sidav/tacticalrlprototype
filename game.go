package main

import (
	"github.com/sidav/golibrl/random/additive_random"
	log2 "tacticalrlprototype/game_log"
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
		CURRENT_MAP.player.performPassbyAttacks()
		for CURRENT_MAP.player.isTimeToAct() && GAME_IS_RUNNING {
			pc.playerControl(&CURRENT_MAP)
		}
		// check if pawns should be removed
		for _, pwn := range CURRENT_MAP.pawns {
			if pwn.hp > pwn.getStaticData().maxhp {
				pwn.hp = pwn.getStaticData().maxhp
			}
			if pwn.isDead() {
				CURRENT_MAP.spitBloodAt(pwn.x, pwn.y)
				CURRENT_MAP.removePawn(pwn)
			}
			if pwn.isTimeToAct() {
				// act for pawns here
				pwn.actAsPawn()
			}
		}
		CURRENT_TURN++
	}
}
