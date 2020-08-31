package crossclient

import (
	"encoding/json"
	"testing"
)

const (
	channelID      = "mychannel"
	orgName        = "Org1"
	orgAdmin       = "Admin"
	ordererOrgName = "OrdererOrg"
	peer1          = "peer0.org1.example.com"
	configPath     = "/opt/gopath/src/github.com/hyperledger/fabric-cross-agent/config/config_e2e.yaml"
)

var (
	ccID = "basic"
)

type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

func TestQuery(t *testing.T) {
	cc, err := NewCrosssClient(channelID, orgName, configPath)
	if err != nil {
		t.Fatalf("Failed to create crossclient:%s", err)
	}
	function := "GetAllAssets"
	args := [][]byte{[]byte("GetALLAssets")}
	res, err := cc.Query(ccID, function, args)
	if err != nil {
		t.Fatal("query err")
	}

	t.Logf("query res : %s", res)
}

func TestInvoke(t *testing.T) {
	cc, err := NewCrosssClient(channelID, orgName, configPath)
	if err != nil {
		t.Fatalf("Failed to create crossclient:%s", err)
	}
	function := "TransferAsset"
	args := [][]byte{[]byte("asset6"), []byte("czbank")}
	res, err := cc.Invoke(ccID, function, args)
	if err != nil {
		t.Fatalf("Failed to invoke :%s", err)
	}

	t.Logf("invoke tx id : %s", res)

	function = "ReadAsset"
	args = [][]byte{[]byte("asset6")}
	queryres, err := cc.Query(ccID, function, args)
	if err != nil {
		t.Fatalf("Faile to query :%s", err)
	}
	asset := Asset{}
	err = json.Unmarshal(queryres, &asset)
	if asset.Owner != "czbank" {
		t.Fatalf("Verify query res failed, query res=%s", queryres)
	}
}
