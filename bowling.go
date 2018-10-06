package bowling

const TEN = 10

type Game struct {
	rolls  []int
	frames int
	length int
}

func (g *Game) Roll(pins int) {
	g.rolls = append(g.rolls, pins)
	g.length = len(g.rolls)
	g.incrementFrame(pins)
}

func (g *Game) Score() int {
	totalRolls := g.length
	if ok, rolls := g.hasTenthFrameBonus(); ok {
		totalRolls = rolls
	}

	score := 0

	for i := 0; i < totalRolls; i++ {
		if g.isStrike(i) {
			score += TEN + g.rolls[i+1] + g.rolls[i+2]
		} else if g.isSpare(i) {
			score += TEN + g.rolls[i+2]
			i++
		} else {
			score += g.rolls[i]
		}
	}

	return score
}

func (g *Game) incrementFrame(pins int) {
	if g.frames != TEN {
		if pins == TEN || g.length%2 == 0 {
			g.frames += 1
		}
	}
}

func (g *Game) isStrike(index int) bool {
	if index+2 < g.length {
		if g.rolls[index] == TEN {
			return true
		}
	}
	return false
}

func (g *Game) isSpare(index int) bool {
	if index+2 < g.length {
		if g.rolls[index]+g.rolls[index+1] == TEN {
			return true
		}
	}

	return false
}

func (g *Game) hasTenthFrameBonus() (bool, int) {
	rollIndex := g.length - 3

	if g.frames == TEN {
		if g.isStrike(rollIndex) || g.isSpare(rollIndex) {
			return true, rollIndex + 1
		}
	}

	return false, 0
}
