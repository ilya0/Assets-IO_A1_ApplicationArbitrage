//Equinix blockchain tool
// v1

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	// ***pb - protocolbuffer
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//this this the users that will be using the
type renewal struct {
	ObjectType string `json:"objecttype"`
	AssetID    string `json:"assetid"`
	CurPrice   string `json:"curprice"`
	PIPercent  string `json:"pipercent"`
	PICap      string `json:"picap"`
	PIPeriod   string `json:"piperiod"`
	NewPrice   string `json:"newprice"`
	State      string `json:"state"`
}

//==============================================================================
// Main
//==============================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoking all the methods for use " + function)

	// Handle different functions
	if function == "sendRenewal" {
		return t.sendRenewal(stub, args)
	} else if function == "updateRenewal" {
		return t.updateRenewal(stub, args)
	} else if function == "getAllRenewals" {
		return t.getAllRenewals(stub, args)
	} else if function == "getAllRenewalsByState" {
		return t.getAllRenewalsByState(stub, args)
	} else if function == "queryRenewal" {
		return t.queryRenewal(stub, args)
	} else if function == "queryHistoryofRenewal" {
		return t.queryHistoryofRenewal(stub, args)
	} else if function == "respondRenewal" {
		return t.respondRenewal(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// Renewal============================================================
// sendRenewal - create a new Renewal
// ============================================================
func (t *SimpleChaincode) sendRenewal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start udpate renewal")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}

	objectType := "renewal"
	renewalId := args[0]
	curPrice := args[1]
	piPercent := args[2]
	piCap := args[3]
	piPeriod := args[4]
	newPrice := args[5]
	state := args[6]

	// ==== Check if renewal already exists ====
	renewalTestAsBytes, err := stub.GetState(renewalId)
	if err != nil {
		return shim.Error("Failed to get renewal: " + err.Error())
	} else if renewalTestAsBytes != nil {
		fmt.Println("This renewal already exists: " + renewalId)
		return shim.Error("This renewal already exists: " + renewalId)
	}

	// ============= Create renewal object and marshal to JSON ====================
	renewal := &renewal{objectType, renewalId, curPrice, piPercent, piCap, piPeriod, newPrice, state}
	renewalAsBytes, err := json.Marshal(renewal)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save renewal to state ===
	err = stub.PutState(renewalId, renewalAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end send renewal ")

	var sendRenewalEventValue []byte
	sendRenewalEventValue = []byte("send renewal with renewal id "+ renewalId)
	stub.SetEvent("sendRenewal", sendRenewalEventValue)

	return shim.Success(renewalAsBytes)

}

// Renewal===============================================
// updateRenewal - read a renewal from chaincode state, if exists update attributes
// ===============================================
func (t *SimpleChaincode) updateRenewal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start udpate renewal")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}

	objectType := "renewal"
	renewalId := args[0]
	curPrice := args[1]
	piPercent := args[2]
	piCap := args[3]
	piPeriod := args[4]
	newPrice := args[5]
	state := args[6]

	// ==== Check if renewal already exists ====
	renewalTestAsBytes, err := stub.GetState(renewalId)
	if err != nil {
		return shim.Error("Failed to get renewal: " + err.Error())
	} else if renewalTestAsBytes == nil {
		fmt.Println("This renewal does not exists: " + renewalId)
		return shim.Error("This renewal does not exists: " + renewalId)
	}

	// ============= Create renewal object and marshal to JSON ====================
	renewal := &renewal{objectType, renewalId, curPrice, piPercent, piCap, piPeriod, newPrice, state}
	renewalAsBytes, err := json.Marshal(renewal)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Update renewal ===
	err = stub.PutState(renewalId, renewalAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end update renewal ")
	return shim.Success(renewalAsBytes)

}

// Renewal==========================================================================
// getAllRenewals - get all renewals
// =================================================================================
func (t *SimpleChaincode) getAllRenewals(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("- start get all renewals ")

	var oneRenewal renewal
	var renewalList []renewal
	var marshalJson []byte

	queryString := fmt.Sprintf("{\"selector\":{\"objecttype\":\"renewal\"}}")
	resultsIterator, _ := stub.GetQueryResult(queryString)
	fmt.Println("queryString : " + queryString)
	fmt.Println(resultsIterator)
	if resultsIterator == nil {
		fmt.Println("No results found for this objectType")
		return shim.Error("No results found for this objectType")
	}
	for resultsIterator.HasNext() {
		fmt.Println("-----------------------hello---------------------------- ")
		resultRow, _ := resultsIterator.Next()
		err := json.Unmarshal([]byte(resultRow.Value), &oneRenewal)
		if err != nil {
			fmt.Println("Error unmarshalling data ", err)
		}

		renewalList = append(renewalList, oneRenewal)
		marshalJson, _ = json.Marshal(renewalList)

	}

	fmt.Println("- end get all renewal ")
	fmt.Println(marshalJson)
	return shim.Success(marshalJson)

}

// Renewal==========================================================================
// getAllRenewalsByState - get all renewals by certain type state
// =================================================================================
func (t *SimpleChaincode) getAllRenewalsByState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("- start get all renewals ")

	var oneRenewal renewal
	var renewalList []renewal
	var marshalJson []byte

	typestate := args[0]
	queryString := "initialized"
	_ = queryString

	if typestate == "accept" {
		queryString = fmt.Sprintf("{\"selector\":{\"state\":\"accept\"}}")
	} else if typestate == "deny" {
		queryString = fmt.Sprintf("{\"selector\":{\"state\":\"deny\"}}")
	} else {
		fmt.Println("This is not a valid state type")
		return shim.Error("Not a valid state type")
	}
	fmt.Println("---- Querystring for getallrenewalsbystate ---")
	fmt.Println(queryString)
	resultsIterator, _ := stub.GetQueryResult(queryString)
	fmt.Println("queryString : " + queryString)
	fmt.Println(resultsIterator)
	if resultsIterator == nil {
		fmt.Println("No results found for this objectType")
		return shim.Error("No results found for this objectType")
	}
	for resultsIterator.HasNext() {
		fmt.Println("-----------------------hello---------------------------- ")
		resultRow, _ := resultsIterator.Next()
		err := json.Unmarshal([]byte(resultRow.Value), &oneRenewal)
		if err != nil {
			fmt.Println("Error unmarshalling data ", err)
		}

		renewalList = append(renewalList, oneRenewal)
		marshalJson, _ = json.Marshal(renewalList)

	}

	fmt.Println("- end get all renewal ")
	fmt.Println(marshalJson)
	return shim.Success(marshalJson)

}

// Renewal==================================================
// queryRenewal - query renewal by ID and return attributes
// ==================================================
func (t *SimpleChaincode) queryRenewal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var RenewalID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Name of the user to query")
	}

	RenewalID = args[0]
	valAsbytes, err := stub.GetState(RenewalID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + RenewalID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"renewal does not exist " + RenewalID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// Renewal============================================================
// queryHistoryofRenewal - query history of renewals by ID and return history of transactions
// ============================================================

func (t *SimpleChaincode) queryHistoryofRenewal(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	intAssetID := args[0]

	fmt.Printf("- start getHistoryForRecord: %s\n", intAssetID)

	resultsIterator, err := stub.GetHistoryForKey(intAssetID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the key/value pair
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON vehiclePart)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// Renewal===============================================
// respondRenewal - Lookup renewal by assetID, if exists change state based on payload - i.e. change state to Accepted or Rejected
// ======================================================
func (t *SimpleChaincode) respondRenewal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	renewalID := args[0]
	state := args[1]

	// ==== Check if renewal already exists ====
	renewalAsBytes, err := stub.GetState(renewalID)
	if err != nil {
		return shim.Error("Failed to get renewal: " + err.Error())
	} else if renewalAsBytes == nil {
		fmt.Println("This renewal does not exists: " + renewalID)
		return shim.Error("This renewal does not exists: " + renewalID)
	}

	//Get data from blockchain and populate it
	renewalData := renewal{}
	err2 := json.Unmarshal(renewalAsBytes, &renewalData)
	if err2 != nil {
		fmt.Println("Unable to unmarshal modify Received PO json::", err)
		return shim.Error(err.Error())
	}
	renewalData.State = state

	renewalJSONasBytes, err := json.Marshal(renewalData)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Renewal to state ===
	err = stub.PutState(renewalID, renewalJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== renewal saved and indexed. Return success ====
	fmt.Println("- end init user")
	return shim.Success(renewalJSONasBytes)
}
