package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const generalAbi = "[{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}," +
	"{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}" +
	"]"

var GeneralAbi, _ = abi.JSON(strings.NewReader(generalAbi))

const balanceOfAbi = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var BalanceOfAbi, _ = abi.JSON(strings.NewReader(balanceOfAbi))

// we abi encode the arg
func ConvertType(arg string, abiinput abi.Type, ds []interface{}) []interface{} {
	if fmt.Sprint(abiinput) == fmt.Sprint("address[]") {
		ds = append(ds, []common.Address{common.HexToAddress(arg)})
	} else if fmt.Sprint(abiinput) == fmt.Sprint("address") {
		ds = append(ds, common.HexToAddress(arg))
	} else if fmt.Sprint(abiinput) == fmt.Sprint("bool") {
		if arg == "1" {
			ds = append(ds, true)
		} else if arg == "0" {
			ds = append(ds, false)
		}
	} else if fmt.Sprint(abiinput) == fmt.Sprint("bool[]") {

		if arg == "1" {
			ds = append(ds, []bool{true})
		} else if arg == "0" {
			ds = append(ds, []bool{false})
		}

	} else if fmt.Sprint(abiinput) == fmt.Sprint("uint256") {
		bn, _ := big.NewInt(0).SetString(arg, 10)
		ds = append(ds, bn)
	}

	return ds
}

// we iterate thru all the output types from the given abi and convert it to a string
func GetOutputString(output []byte, outputTypes abi.Arguments, methodname string, abiToTest abi.ABI) string {
	res := ""
	for _, abiinput := range outputTypes {

		if fmt.Sprint(abiinput.Type) == fmt.Sprint("uint256[]") {
			var intres []*big.Int
			abiToTest.UnpackIntoInterface(&intres, methodname, output)
			for i := 0; i < len(intres); i++ {
				res = res + intres[i].String() + "|"
			}
		} else if fmt.Sprint(abiinput.Type) == fmt.Sprint("uint256") {
			var intres *big.Int
			abiToTest.UnpackIntoInterface(&intres, methodname, output)
			res = res + intres.String() + "|"
		}
	}

	return res
}

func AbiParser(client *ethclient.Client,
	contract common.Address,
	methodname string,
	args []string,
	abiToTest abi.ABI) (string, error) {
	//mybal := tiblockchain.GetTokenBalance(client, token, from)
	ds := []interface{}{}

	for _, _m := range abiToTest.Methods {
		if fmt.Sprint(_m.Name) == methodname {
			if len(args) == len(_m.Inputs) {
				for i := 0; i < len(_m.Inputs); i++ {
					ds = ConvertType(args[i], _m.Inputs[i].Type, ds)
				}
				if len(ds) > 0 {
					swapData, _ := abiToTest.Pack(_m.Name, ds...)
					if _m.StateMutability == "view" {
						msgz := ethereum.CallMsg{
							To:   &contract,
							Data: swapData,
						}
						outputz, _ := client.CallContract(context.Background(), msgz, nil)
						res := GetOutputString(outputz, _m.Outputs, _m.Name, abiToTest)
						return res, nil

					}
				}

				break
			}
		}

	}
	return "", errors.New("NOT FOUND")
}
