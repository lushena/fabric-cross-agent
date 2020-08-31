package crossclient

import "testing"

const (
	channelID      = "mychannel"
	orgName        = "Org1"
	orgAdmin       = "Admin"
	ordererOrgName = "OrdererOrg"
	peer1          = "peer0.org1.example.com"
)

var (
	ccID = "basic"
)

func TestQuery(t *testing.T) {
	cc, err := NewCrosssClient()
	if err != nil {
		t.Fatalf("Failed to create crossclient:%s", err)
	}
	function := "GetAllAssets"
	args := [][]byte{[]byte("GetALLAssets")}
	res, err := cc.Query(ccID, function, channelID, orgName, args)
	if err != nil {
		t.Fatal("query err")
	}

	t.Logf("query res : %s", res)
}
