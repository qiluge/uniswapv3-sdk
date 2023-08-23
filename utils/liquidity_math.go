package utils

import (
	"math/big"
)

func AddDelta(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}
