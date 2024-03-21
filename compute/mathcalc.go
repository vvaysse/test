package compute

import "math/big"

func MultiplyXByY(x *big.Int, yFloat *big.Float) *big.Int {
	// Convert the big.Int to a big.Float
	xFloat := new(big.Float)
	xFloat.SetInt(new(big.Int).SetBytes(x.Bytes()))

	// Multiply xFloat by yFloat and store the result in a new big.Float
	zFloat := new(big.Float)
	zFloat.Mul(xFloat, yFloat)
	zFloatInt, _ := zFloat.Int(nil)
	zFloatBytes := zFloatInt.Bytes()

	// Convert the result to a big.Int
	zInt := new(big.Int)
	zInt.SetBytes(zFloatBytes)

	// Return the result
	return zInt
}
