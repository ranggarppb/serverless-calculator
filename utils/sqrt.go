package utils

import (
	"github.com/shopspring/decimal"
)

func Sqrt(d decimal.Decimal) decimal.Decimal {
	// s, _ := d.SqrtRound(int32(DivisionPrecision))
	res, _ := sqrtRound(d, int32(31))
	return res
}

func sqrtRound(d decimal.Decimal, precision int32) (decimal.Decimal, bool) {
	const sqrtMaxIter = 10000

	if d.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero, false
	}
	cutoff := decimal.New(1, -precision)
	lo := decimal.Zero
	hi := d
	var mid decimal.Decimal
	for i := 0; i < sqrtMaxIter; i++ {
		//mid = (lo+hi)/2;
		mid = lo.Add(hi).DivRound(decimal.NewFromInt(2), precision)
		if mid.Mul(mid).Sub(d).Abs().LessThan(cutoff) {
			return mid, true
		}
		if mid.Mul(mid).GreaterThan(d) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return mid, false
}
