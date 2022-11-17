// Code generated by codegen; DO NOT EDIT.
package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

type KeyvaultClient struct {
	KeysClient        KeysClient
	ManagedHsmsClient ManagedHsmsClient
	SecretsClient     SecretsClient
	VaultsClient      VaultsClient
}

func NewKeyvaultClient(subscriptionID string, credentials azcore.TokenCredential, options *arm.ClientOptions) (*KeyvaultClient, error) {
	var client KeyvaultClient
	var err error

	client.KeysClient, err = armkeyvault.NewKeysClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	client.ManagedHsmsClient, err = armkeyvault.NewManagedHsmsClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	client.SecretsClient, err = armkeyvault.NewSecretsClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	client.VaultsClient, err = armkeyvault.NewVaultsClient(subscriptionID, credentials, options)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
