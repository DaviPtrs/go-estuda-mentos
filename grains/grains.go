package grains

import (
	"errors"
	"math"
	"math/big"
)

const ChessSquares int = 64
const qFactor int = 2

func Square(number int) (uint64, error) {
	if number < 1 || number > ChessSquares {
		return 0, errors.New("invalid number")
	}
	return uint64(1 * math.Pow(float64(qFactor), float64(number-1))), nil
}

func Total() uint64 {
	num := big.NewInt(int64(ChessSquares))
	factor := big.NewInt(int64(qFactor))
	return num.Exp(factor, num, nil).Sub(num, big.NewInt(1)).Uint64()
}
