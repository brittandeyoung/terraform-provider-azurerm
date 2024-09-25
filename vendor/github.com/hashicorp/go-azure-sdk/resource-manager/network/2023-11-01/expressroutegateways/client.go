package expressroutegateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*ExpressRouteGatewaysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "expressroutegateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteGatewaysClient: %+v", err)
	}

	return &ExpressRouteGatewaysClient{
		Client: client,
	}, nil
}
