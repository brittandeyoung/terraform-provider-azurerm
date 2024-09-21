package managementgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionsGetSubscriptionsUnderManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionUnderManagementGroup
}

type SubscriptionsGetSubscriptionsUnderManagementGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubscriptionUnderManagementGroup
}

type SubscriptionsGetSubscriptionsUnderManagementGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SubscriptionsGetSubscriptionsUnderManagementGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SubscriptionsGetSubscriptionsUnderManagementGroup ...
func (c ManagementGroupsClient) SubscriptionsGetSubscriptionsUnderManagementGroup(ctx context.Context, id commonids.ManagementGroupId) (result SubscriptionsGetSubscriptionsUnderManagementGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SubscriptionsGetSubscriptionsUnderManagementGroupCustomPager{},
		Path:       fmt.Sprintf("%s/subscriptions", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]SubscriptionUnderManagementGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SubscriptionsGetSubscriptionsUnderManagementGroupComplete retrieves all the results into a single object
func (c ManagementGroupsClient) SubscriptionsGetSubscriptionsUnderManagementGroupComplete(ctx context.Context, id commonids.ManagementGroupId) (SubscriptionsGetSubscriptionsUnderManagementGroupCompleteResult, error) {
	return c.SubscriptionsGetSubscriptionsUnderManagementGroupCompleteMatchingPredicate(ctx, id, SubscriptionUnderManagementGroupOperationPredicate{})
}

// SubscriptionsGetSubscriptionsUnderManagementGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagementGroupsClient) SubscriptionsGetSubscriptionsUnderManagementGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, predicate SubscriptionUnderManagementGroupOperationPredicate) (result SubscriptionsGetSubscriptionsUnderManagementGroupCompleteResult, err error) {
	items := make([]SubscriptionUnderManagementGroup, 0)

	resp, err := c.SubscriptionsGetSubscriptionsUnderManagementGroup(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = SubscriptionsGetSubscriptionsUnderManagementGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
