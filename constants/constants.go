package constants

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

const PoolInitCodeHash = "0xe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54"

var (
	FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressZero    = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// The default factory enabled fee amounts, denominated in hundredths of bips.
type FeeAmount uint64

const (
	FeeLowest FeeAmount = 100
	FeeLow    FeeAmount = 500
	FeeMedium FeeAmount = 3000
	FeeHigh   FeeAmount = 10000

	FeeMax FeeAmount = 1000000

	FeePancake25 = 2500

	FeeMin    = 50
	Fee250    = 250
	Fee25     = 25
	Fee5      = 5
	Fee10     = 10
	FeeMinMin = 2

	Fee400 = 400
	Fee450 = 450
)

// The default factory tick spacings by fee amount.
var TickSpacings = map[FeeAmount]int{
	FeeMinMin:    1,
	FeeMin:       1,
	FeeLowest:    1,
	Fee5:         1,
	Fee10:        1,
	Fee25:        1,
	Fee250:       5,
	FeeLow:       10,
	Fee400:       10,
	Fee450:       10,
	FeeMedium:    60,
	FeeHigh:      200,
	FeePancake25: 50,
}

var (
	NegativeOne = big.NewInt(-1)
	Zero        = big.NewInt(0)
	One         = big.NewInt(1)

	// used in liquidity amount math
	Q96  = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
	Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)

	PercentZero = entities.NewFraction(big.NewInt(0), big.NewInt(1))
)
