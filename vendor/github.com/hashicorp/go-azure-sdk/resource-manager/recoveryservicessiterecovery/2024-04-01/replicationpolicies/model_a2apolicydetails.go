package replicationpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyProviderSpecificDetails = A2APolicyDetails{}

type A2APolicyDetails struct {
	AppConsistentFrequencyInMinutes   *int64  `json:"appConsistentFrequencyInMinutes,omitempty"`
	CrashConsistentFrequencyInMinutes *int64  `json:"crashConsistentFrequencyInMinutes,omitempty"`
	MultiVMSyncStatus                 *string `json:"multiVmSyncStatus,omitempty"`
	RecoveryPointHistory              *int64  `json:"recoveryPointHistory,omitempty"`
	RecoveryPointThresholdInMinutes   *int64  `json:"recoveryPointThresholdInMinutes,omitempty"`

	// Fields inherited from PolicyProviderSpecificDetails

	InstanceType string `json:"instanceType"`
}

func (s A2APolicyDetails) PolicyProviderSpecificDetails() BasePolicyProviderSpecificDetailsImpl {
	return BasePolicyProviderSpecificDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = A2APolicyDetails{}

func (s A2APolicyDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2APolicyDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2APolicyDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2APolicyDetails: %+v", err)
	}

	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2APolicyDetails: %+v", err)
	}

	return encoded, nil
}
