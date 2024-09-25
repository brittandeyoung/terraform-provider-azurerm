package fleets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FleetsClient struct {
	Client *resourcemanager.Client
}

func NewFleetsClientWithBaseURI(sdkApi sdkEnv.Api) (*FleetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "fleets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FleetsClient: %+v", err)
	}

	return &FleetsClient{
		Client: client,
	}, nil
}
