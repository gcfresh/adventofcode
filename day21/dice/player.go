package dice

type Player struct {
	ID      int
	Start   int
	Current int
	Score   int
}

func (p *Player) Move(num int) {
	p.Current = (p.Current + num) % 10
	if p.Current == 0 {
		p.Current = 10
	}
	p.Score = p.Score + p.Current
}
