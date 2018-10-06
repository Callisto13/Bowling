package bowling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Callisto13/bowling"
)

var _ = Describe("Bowling", func() {
	var game bowling.Game

	BeforeEach(func() {
		game = bowling.Game{}
	})

	Describe("Score", func() {
		Context("should be able to call score at any time", func() {
			It("such as after one roll", func() {
				game.Roll(5)
				Expect(game.Score()).To(Equal(5))
			})

			It("such as after one strike", func() {
				game.Roll(10)
				Expect(game.Score()).To(Equal(10))
			})

			It("such as after one spare", func() {
				game.Roll(5)
				game.Roll(5)
				Expect(game.Score()).To(Equal(10))
			})
		})
	})

	Context("gutter game", func() {
		It("should score 0 for a gutter game, all rolls are 0", func() {
			g := rollMany(game, 20, 0)
			Expect(g.Score()).To(Equal(0))
		})
	})

	Context("given all rolls hit only one pin", func() {
		It("should score 20", func() {
			g := rollMany(game, 20, 1)
			Expect(g.Score()).To(Equal(20))
		})
	})

	Context("given a game with a spare", func() {
		It("should score 20 given the first 3 rolls hit 5 pins", func() {
			game.Roll(5)
			game.Roll(5)
			game.Roll(5)

			Expect(game.Score()).To(Equal(20))
		})
	})

	Context("given a game with only one strike", func() {
		It("should score 20 given a strike followed by two rolls hitting 2 & 3 pins", func() {
			game.Roll(10)
			game.Roll(2)
			game.Roll(3)
			Expect(game.Score()).To(Equal(20))
		})
	})

	Context("perfect game", func() {
		It("should score 300 for 12 strikes in a row", func() {
			g := rollMany(game, 12, 10)
			Expect(g.Score()).To(Equal(300))
		})
	})

	Context("game with all scoring variations including tenth frame", func() {
		It("should score 110", func() {
			// frame 1, score: 9
			game.Roll(7)
			game.Roll(2)
			// frame 2, score: 16
			game.Roll(6)
			game.Roll(1)
			// frame 3, score: 26 + 3 = 29
			game.Roll(5)
			game.Roll(5)
			// frame 4, score: 36
			game.Roll(3)
			game.Roll(4)
			// frame 5, score: 46 + 10 = 56
			game.Roll(5)
			game.Roll(5)
			// frame 6, score: 66 + 5 + 3 = 74
			game.Roll(10)
			// frame 7, score: 82
			game.Roll(5)
			game.Roll(3)
			// frame 8, score: 87
			game.Roll(5)
			game.Roll(0)
			// frame 9, score: 95
			game.Roll(6)
			game.Roll(2)
			// frame 10, score: 105 + 5 = 110
			game.Roll(7)
			game.Roll(3)
			game.Roll(5)
			Expect(game.Score()).To(Equal(110))
		})
	})
})

func rollMany(game bowling.Game, rolls, pins int) bowling.Game {
	for i := 0; i < rolls; i++ {
		game.Roll(pins)
	}

	return game
}
