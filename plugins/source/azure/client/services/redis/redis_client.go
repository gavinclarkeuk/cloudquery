// Code generated by codegen; DO NOT EDIT.
package redis

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

type RedisClient struct {
	Client Client
}

func NewRedisClient(subscriptionID string, credentials azcore.TokenCredential, options *arm.ClientOptions) (*RedisClient, error) {
	var client RedisClient
	var err error

	client.Client, err = armredis.NewClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
