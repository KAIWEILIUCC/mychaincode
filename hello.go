package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strconv"
)

type SmartContract struct {
	contractapi.Contract
}


func (t *SmartContract)init(ctx contractapi.TransactionContextInterface, name string, age int, sex string) error{
	fmt.Println("Init the ledger...")
	var err error

	fmt.Printf("name:%s, age:%d, sex:%s\n", name, age, sex)

	err = ctx.GetStub().PutState("name", []byte(name))
	if err != nil{
		return err
	}

	err = ctx.GetStub().PutState("age", []byte(strconv.Itoa(age)))
	if err != nil{
		return err
	}

	err = ctx.GetStub().PutState("sex", []byte(sex))
	if err != nil{
		return err
	}

	return nil
}

func (t *SmartContract)query(ctx contractapi.TransactionContextInterface, term string) (string, error){
	var err error
	termByteArray,err := ctx.GetStub().GetState(term)

	if err != nil{
		jsonResp := "{\"Error\":\"Failed to get state for " + term + "\"}"
		return "", errors.New(jsonResp)
	}

	if termByteArray == nil{
		jsonResp := "{\"Error\":\"Nil amount for " + term + "\"}"
		return "", errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + term + "\",\"Amount\":\"" + string(termByteArray) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return string(termByteArray), nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting this chaincode: %s", err)
	}
}
