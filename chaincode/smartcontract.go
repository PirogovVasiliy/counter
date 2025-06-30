package chaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type CounterContract struct{ contractapi.Contract }

type Counter struct {
	Value int `json:"value"`
}

const stateCounter string = "stateCounter"

func (cc *CounterContract) CreateCounter(ctx contractapi.TransactionContextInterface, initValue int) error {
	counter := Counter{Value: initValue}
	counterJSON, err := json.Marshal(counter)
	if err != nil {
		return err
	}

	ctx.GetStub().SetEvent("CreateCounter", counterJSON)

	return ctx.GetStub().PutState(stateCounter, counterJSON)
}

func (cc *CounterContract) ReadCounter(ctx contractapi.TransactionContextInterface) (*Counter, error) {
	counterJSON, err := ctx.GetStub().GetState(stateCounter)
	if err != nil {
		return nil, err
	}

	var counter Counter
	err = json.Unmarshal(counterJSON, &counter)
	if err != nil {
		return nil, err
	}

	return &counter, nil
}

func (cc *CounterContract) IncrimentCounter(ctx contractapi.TransactionContextInterface, value int) error {
	counterJSON, err := ctx.GetStub().GetState(stateCounter)
	if err != nil {
		return err
	}

	var counter Counter
	err = json.Unmarshal(counterJSON, &counter)
	if err != nil {
		return err
	}

	counter.Value += value

	counterJSON, err = json.Marshal(counter)
	if err != nil {
		return err
	}

	ctx.GetStub().SetEvent("IncrimentCounter", counterJSON)

	return ctx.GetStub().PutState(stateCounter, counterJSON)
}

func (cc *CounterContract) MinusCounter(ctx contractapi.TransactionContextInterface) error {
	counterJSON, err := ctx.GetStub().GetState(stateCounter)
	if err != nil {
		return err
	}

	var counter Counter
	err = json.Unmarshal(counterJSON, &counter)
	if err != nil {
		return err
	}

	counter.Value -= 1

	counterJSON, err = json.Marshal(counter)
	if err != nil {
		return err
	}

	ctx.GetStub().SetEvent("MinusCounter", counterJSON)

	return ctx.GetStub().PutState(stateCounter, counterJSON)
}
