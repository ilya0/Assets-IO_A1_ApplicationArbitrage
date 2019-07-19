package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	// "strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestSendRenewal(t *testing.T) {

	fmt.Println("----------Entered into TestSendRenewal-----------")

	var renewalId = "10"

	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleChaincode))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	args := [][]byte{
		[]byte("sendRenewal"),
		[]byte(renewalId),
		[]byte("23"),
		[]byte("60"),
		[]byte("100"),
		[]byte("2"),
		[]byte("25"),
		[]byte("approved")}
	invokeResult := stub.MockInvoke("12345", args)
	if invokeResult.Status != shim.OK {
		fmt.Println("--------------Error----------")
		fmt.Println(invokeResult)
		t.Fatalf("Expected unauthorized user error to be returned")
	}

	getResultAsBytes, err := stub.GetState(renewalId)
	if err != nil {
		t.Errorf("Failed to get state for " + renewalId)
	} else if getResultAsBytes == nil {
		t.Errorf("renewal does not exist: " + renewalId)
	}

	renewalResult := renewal{"renewal", renewalId, "23", "60", "100", "2", "25", "approved"}

	//convert getResultAsBytes to []string
	resultJSON := renewal{}
	error := json.Unmarshal(getResultAsBytes, &resultJSON)
	if error != nil {
		t.Errorf("Error unmarshalling getPartsByAssembler JSON result")
	}
	fmt.Println(resultJSON)
	fmt.Println(renewalResult)

	isEqual := reflect.DeepEqual(resultJSON, renewalResult)
	if !isEqual {
		t.Fatalf("Unexpected response received")
	}
}

func TestUpdateRenewal(t *testing.T) {

	fmt.Println("----------Entered into Test Update Renewal-----------")

	var renewalId = "10"

	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleChaincode))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	//=================== Call the update method now ====================
	args2 := [][]byte{
		[]byte("updateRenewal"),
		[]byte(renewalId),
		[]byte("27"),
		[]byte("65"),
		[]byte("102"),
		[]byte("4"),
		[]byte("50"),
		[]byte("rejected")}
	invokeResult2 := stub.MockInvoke("123455", args2)
	if invokeResult2.Status != shim.OK {
		fmt.Println("--------------Error----------")
		fmt.Println(invokeResult2)
		t.Fatalf("Expected unauthorized user error to be returned")
	}

	getResultAsBytes, err := stub.GetState(renewalId)
	if err != nil {
		t.Errorf("Failed to get state for " + renewalId)
	} else if getResultAsBytes == nil {
		t.Errorf("renewal does not exist: " + renewalId)
	}

	renewalResult := renewal{"renewal", renewalId, "27", "65", "102", "4", "50", "rejected"}

	//convert getResultAsBytes to []string
	resultJSON := renewal{}
	error := json.Unmarshal(getResultAsBytes, &resultJSON)
	if error != nil {
		t.Errorf("Error unmarshalling getPartsByAssembler JSON result")
	}
	fmt.Println(resultJSON)
	fmt.Println(renewalResult)

	isEqual := reflect.DeepEqual(resultJSON, renewalResult)
	if !isEqual {
		t.Fatalf("Unexpected response received")
	}

}

func TestGetAllRenewals(t *testing.T) {
	fmt.Println("----------Entered into Test Update Renewal-----------")

	var renewalId = "11"

	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleChaincode))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	//=================== Call the create method again, so we have 2 renewals in blockchain ====================
	args3 := [][]byte{
		[]byte("sendRenewal"),
		[]byte("11"),
		[]byte("23"),
		[]byte("60"),
		[]byte("100"),
		[]byte("2"),
		[]byte("25"),
		[]byte("approved")}
	invokeResult3 := stub.MockInvoke("12345", args3)
	if invokeResult3.Status != shim.OK {
		fmt.Println("--------------Error----------")
		fmt.Println(invokeResult3)
		t.Fatalf("Expected unauthorized user error to be returned")
	}

	args := [][]byte{
		[]byte("sendRenewal"),
		[]byte("12"),
		[]byte("234"),
		[]byte("600"),
		[]byte("1000"),
		[]byte("20"),
		[]byte("250"),
		[]byte("rejected")}
	invokeResult := stub.MockInvoke("12345", args)
	if invokeResult.Status != shim.OK {
		fmt.Println("--------------Error----------")
		fmt.Println(invokeResult)
		t.Fatalf("Expected unauthorized user error to be returned")
	}

	getResultAsBytes, err := stub.GetState(renewalId)
	if err != nil {
		t.Errorf("Failed to get state for " + renewalId)
	} else if getResultAsBytes == nil {
		t.Errorf("renewal does not exist: " + renewalId)
	}
	fmt.Println(getResultAsBytes)

	args2 := [][]byte{[]byte("getAllRenewals")}
	invokeResult2 := stub.MockInvoke("12345", args2)
	if invokeResult2.Status != shim.OK {
		fmt.Println("--------------Error----------")
		fmt.Println(invokeResult2)
		t.Fatalf("Expected unauthorized user error to be returned")
	}

	fmt.Println(invokeResult2)
}

/// ++++TestRespondRenewals test -
// func TestRespondRenewals(t *testing.T) {
// 	fmt.Println("----------Entered into Test Respond Renewal-----------")

// 	var renewalId = "1"

// 	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleChaincode))
// 	if stub == nil {
// 		t.Fatalf("MockStub creation failed")
// 	}

// 	//=================== Call the create method again, so we have 2 renewals in blockchain ====================
// 	args3 := [][]byte{
// 		[]byte("sendRenewal"),
// 		[]byte("1"),
// 		[]byte("23"),
// 		[]byte("60"),
// 		[]byte("100"),
// 		[]byte("2"),
// 		[]byte("25"),
// 		[]byte("approved")}
// 	invokeResult3 := stub.MockInvoke("12345", args3)
// 	if invokeResult3.Status != shim.OK {
// 		fmt.Println("--------------Error----------")
// 		fmt.Println(invokeResult3)
// 		t.Fatalf("Expected unauthorized user error to be returned")
// 	}

// 	args := [][]byte{
// 		[]byte("sendRenewal"),
// 		[]byte("12"),
// 		[]byte("234"),
// 		[]byte("600"),
// 		[]byte("1000"),
// 		[]byte("20"),
// 		[]byte("250"),
// 		[]byte("rejected")}
// 	invokeResult := stub.MockInvoke("12345", args)
// 	if invokeResult.Status != shim.OK {
// 		fmt.Println("--------------Error----------")
// 		fmt.Println(invokeResult)
// 		t.Fatalf("Expected unauthorized user error to be returned")
// 	}

// 	getResultAsBytes, err := stub.GetState(renewalId)
// 	if err != nil {
// 		t.Errorf("Failed to get state for " + renewalId)
// 	} else if getResultAsBytes == nil {
// 		t.Errorf("renewal does not exist: " + renewalId)
// 	}
// 	fmt.Println(getResultAsBytes)

// 	args2 := [][]byte{[]byte("getAllRenewals")}
// 	invokeResult2 := stub.MockInvoke("12345", args2)
// 	if invokeResult2.Status != shim.OK {
// 		fmt.Println("--------------Error----------")
// 		fmt.Println(invokeResult2)
// 		t.Fatalf("Expected unauthorized user error to be returned")
// 	}

// 	fmt.Println(invokeResult2)
// }
