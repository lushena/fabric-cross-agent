package crossclient

import (
	"fmt"

	"github.com/hyperledger/fabric-cross-agent/integration"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type CrossClient struct {
	FabSDK *fabsdk.FabricSDK
}

func NewCrosssClient(sdkOpts ...fabsdk.Option) (*CrossClient, error) {
	configPath := "/opt/gopath/src/github.com/hyperledger/fabric-cross-agent/config/config_e2e.yaml"
	configOpt := config.FromFile(configPath)
	if integration.IsLocal() {
		//If it is a local test then add entity mapping to config backend to parse URLs
		configOpt = integration.AddLocalEntityMapping(configOpt)
	}
	fsdk, err := fabsdk.New(configOpt, sdkOpts...)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new SDK: %s", err)
	}
	c := &CrossClient{}
	c.FabSDK = fsdk
	return c, nil
}

func (cc *CrossClient) Query(ccID, function, channelID, orgName string, args [][]byte, targetEndpoints ...string) ([]byte, error) {
	clientChannelContext := cc.FabSDK.ChannelContext(channelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg(orgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("channel client error:%s", err)
	}
	response, err := client.Query(channel.Request{ChaincodeID: ccID, Fcn: function, Args: args},
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(targetEndpoints...),
	)
	if err != nil {
		return nil, fmt.Errorf("query ccid:%s,function:%s, channel:%s, orgname:%s err:%s", ccID, function, channelID, orgName, err)
	}
	return response.Payload, err
}
