package blockchain

import (
	"bcback/compute"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBaseFee(c *ethclient.Client) *big.Int {
	for {
		bn, errone := c.BlockNumber(context.Background())
		if errone == nil {
			blocknb := new(big.Int).SetUint64(bn)

			if blocknb.Cmp(big.NewInt(0)) <= 0 {
				return big.NewInt(105659034679) // arbitrary base fee based on block 17188130
			}

			blk, errtwo := c.BlockByNumber(context.Background(), blocknb)
			if errtwo == nil {
				return compute.MultiplyXByY(blk.Header().BaseFee, big.NewFloat(1.20))
			}
		}
	}

}
