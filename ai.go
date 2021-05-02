package main

func (p *pawn) actAsPawn() {
	if p.ai == nil {
		log.AppendMessage("no ai error")
	}
}

type ai struct {}
