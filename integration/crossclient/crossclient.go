package crossclient

import (
	"fmt"

	"github.com/hyperledger/fabric-cross-agent/integration"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type CrossClient struct {
	FabSDK        *fabsdk.FabricSDK
	ChannelClient *channel.Client
}

func NewCrosssClient(channelID, orgName, configPath string, sdkOpts ...fabsdk.Option) (*CrossClient, error) {

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

	clientChannelContext := fsdk.ChannelContext(channelID, fabsdk.WithUser("Admin"), fabsdk.WithOrg(orgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("channel client error:%s", err)
	}
	c.ChannelClient = client
	return c, nil
}

func (cc *CrossClient) Query(ccID, function string, args [][]byte, targetEndpoints ...string) ([]byte, error) {

	response, err := cc.ChannelClient.Query(channel.Request{ChaincodeID: ccID, Fcn: function, Args: args},
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(targetEndpoints...),
	)
	if err != nil {
		return nil, fmt.Errorf("query ccid:%s,function:%s, err:%s", ccID, function, err)
	}
	return response.Payload, err
}

func (cc *CrossClient) Invoke(ccID, function string, args [][]byte) (fab.TransactionID, error) {
	response, err := cc.ChannelClient.Execute(channel.Request{ChaincodeID: ccID, Fcn: function, Args: args}, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return "", fmt.Errorf("invoke ccid:%s,err:%s", ccID, err)
	}

	return response.TransactionID, nil
}
