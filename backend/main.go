package main

import (
	"bcback/blockchain"
	"bcback/httpOut"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func verifyfrom(client *ethclient.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		adr := r.URL.Query().Get("adr")
		addressWhoSigned := common.HexToAddress(adr)
		contract := r.URL.Query().Get("ca")
		network := r.URL.Query().Get("network")

		//fmt.Println("RAW MSG: ", rawmsg, " ADDRESS:", adr, "SIGNED HASH: ", signhash)

		CaToVeferify := common.HexToAddress(contract)
		emptyargs := []string{addressWhoSigned.Hex()}
		res, err := blockchain.AbiParser(client, CaToVeferify, "balanceOf", emptyargs, blockchain.BalanceOfAbi)
		if err == nil {
			if len(res) > 0 && res != "0|" {
				fmt.Println(res, network)

				httpOut.TextResp(w, "True")
				//fmt.Println("TRUE")
				return
			}
		}

		httpOut.TextResp(w, "False")
	}

}
func main() {

	// CHANGE THE POLYGON RPC IF NEEDED

	ipc := "https://polygon.rpc.blxrbdn.com"

	c, cerr := ethclient.Dial(ipc)
	if cerr != nil {
		log.Fatal("Bad RPC address")
	}

	http.HandleFunc("/verifyfrom", verifyfrom(c))

	http.HandleFunc("/sign", func(w http.ResponseWriter, r *http.Request) {

		rawmsg := r.URL.Query().Get("msg")
		adr := r.URL.Query().Get("adr")
		signhash := r.URL.Query().Get("sig")

		addressWhoSigned := blockchain.GetAddressFromSign(rawmsg, signhash)

		//fmt.Println("RAW MSG: ", rawmsg, " ADDRESS:", adr, "SIGNED HASH: ", signhash)
		if addressWhoSigned == common.HexToAddress(adr) {
			httpOut.TextResp(w, "True")
			//fmt.Println("TRUE")
			return

		} else {
			//fmt.Println("FALSE")
			httpOut.TextResp(w, "False")
		}

		// }
	})

	err := http.ListenAndServe(":4365", nil)
	fmt.Println(err)
}
