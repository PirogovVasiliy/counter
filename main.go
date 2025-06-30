package main

import (
	"counter/chaincode"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	counterContr, err := contractapi.NewChaincode(&chaincode.CounterContract{})
	if err != nil {
		log.Panicf("Error creating counter chaincode: %v", err)
	}

	if err := counterContr.Start(); err != nil {
		log.Panicf("Error starting counter chaincode: %v", err)
	}
}
