package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetAddressFromSign(message string, signature string) common.Address {
	messageWithPrefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(messageWithPrefix))
	signatureBytes := common.FromHex(signature)
	if len(signatureBytes) != 65 {
		return common.HexToAddress("0x")
	}
	signatureBytes[64] -= 27 // Transform V from 27/28 to 0/1 according to Ethereum's yellow paper
	pubkey, err := crypto.SigToPub(messageHash.Bytes(), signatureBytes)
	if err != nil {
		return common.HexToAddress("0x")
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubkey)
	return recoveredAddr

}
