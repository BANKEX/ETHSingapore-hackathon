package events

import (
	"../../config"
	"../../ethereum/plasmacontract"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func ListenDeposit() {

	client, err := ethclient.Dial(config.GetVerifier().GethHost)
	if err != nil {
		log.Println(err)
	}

	contractAddress := common.HexToAddress(config.GetVerifier().PlasmaContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {

	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {

	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			Sig := []byte("CoinDeposited(address,uint256)")
			SigHash := crypto.Keccak256Hash(Sig)
			switch vLog.Topics[0].Hex() {
			case SigHash.Hex():
				event := struct {
					Who    common.Address
					Amount *big.Int
				}{}
				err := contractAbi.Unpack(&event, "CoinDeposited", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}


				// fmt.Println("------------------------------")
				// fmt.Println("Triggered event: CoinDeposited")

				// fmt.Print("Who: ")
				// fmt.Println(common.HexToAddress(vLog.Topics[1].Hex()).String())

				// fmt.Print("Sum in wei: ")
				// fmt.Println(event.Amount.String())
				// fmt.Println("------------------------------")

			}
		}
	}
}
