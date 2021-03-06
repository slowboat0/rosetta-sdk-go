// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestExpectedOperation(t *testing.T) {
	var tests = map[string]struct {
		intent   *types.Operation
		observed *types.Operation

		err bool
	}{
		"simple match": {
			intent: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 1,
				},
				Type: "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			observed: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 3,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index: 2,
					},
				},
				Status: "success",
				Type:   "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
		},
		"account mismatch": {
			intent: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 1,
				},
				Type: "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			observed: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 3,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index: 2,
					},
				},
				Status: "success",
				Type:   "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr2",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			err: true,
		},
		"amount mismatch": {
			intent: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 1,
				},
				Type: "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			observed: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 3,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index: 2,
					},
				},
				Status: "success",
				Type:   "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "150",
				},
			},
			err: true,
		},
		"type mismatch": {
			intent: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 1,
				},
				Type: "transfer",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			observed: &types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index: 3,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index: 2,
					},
				},
				Status: "success",
				Type:   "reward",
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Amount: &types.Amount{
					Value: "100",
				},
			},
			err: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := ExpectedOperation(test.intent, test.observed)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExpectedOperations(t *testing.T) {
	var tests = map[string]struct {
		intent         []*types.Operation
		observed       []*types.Operation
		errExtra       bool
		confirmSuccess bool

		err bool
	}{
		"simple match": {
			intent: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type: "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 5,
					},
					Type: "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
			},
			observed: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 2,
					},
					Status: "success",
					Type:   "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 3,
					},
					RelatedOperations: []*types.OperationIdentifier{
						{
							Index: 2,
						},
					},
					Status: "success",
					Type:   "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
		},
		"simple match (confirm success)": {
			intent: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type: "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 5,
					},
					Type: "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
			},
			observed: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 2,
					},
					Status: "success",
					Type:   "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 3,
					},
					RelatedOperations: []*types.OperationIdentifier{
						{
							Index: 2,
						},
					},
					Status: "success",
					Type:   "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
			confirmSuccess: true,
		},
		"simple match (confirm success) err": {
			intent: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type: "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 5,
					},
					Type: "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
			},
			observed: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 2,
					},
					Status: "success",
					Type:   "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 3,
					},
					RelatedOperations: []*types.OperationIdentifier{
						{
							Index: 2,
						},
					},
					Status: "failure",
					Type:   "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
			confirmSuccess: true,
			err:            true,
		},
		"err extra": {
			intent: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type: "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
			observed: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 2,
					},
					Status: "success",
					Type:   "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 3,
					},
					RelatedOperations: []*types.OperationIdentifier{
						{
							Index: 2,
						},
					},
					Status: "success",
					Type:   "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
			errExtra: true,
			err:      true,
		},
		"missing match": {
			intent: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 1,
					},
					Type: "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 5,
					},
					Type: "fee",
					Account: &types.AccountIdentifier{
						Address: "addr2",
					},
					Amount: &types.Amount{
						Value: "50",
					},
				},
			},
			observed: []*types.Operation{
				{
					OperationIdentifier: &types.OperationIdentifier{
						Index: 3,
					},
					RelatedOperations: []*types.OperationIdentifier{
						{
							Index: 2,
						},
					},
					Status: "success",
					Type:   "transfer",
					Account: &types.AccountIdentifier{
						Address: "addr1",
					},
					Amount: &types.Amount{
						Value: "100",
					},
				},
			},
			err: true,
		},
	}

	for name, test := range tests {
		asserter, err := simpleAsserterConfiguration([]*types.OperationStatus{
			{
				Status:     "success",
				Successful: true,
			},
			{
				Status:     "failure",
				Successful: false,
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, asserter)

		parser := New(asserter, nil)

		t.Run(name, func(t *testing.T) {
			err := parser.ExpectedOperations(
				test.intent,
				test.observed,
				test.errExtra,
				test.confirmSuccess,
			)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExpectedSigners(t *testing.T) {
	var tests = map[string]struct {
		intent   []*types.SigningPayload
		observed []string

		err bool
	}{
		"simple match": {
			intent: []*types.SigningPayload{
				{
					Address: "addr1",
				},
				{
					Address: "addr2",
				},
				{
					Address: "addr2",
				},
			},
			observed: []string{
				"addr1",
				"addr2",
			},
		},
		"duplicate observed signers": {
			intent: []*types.SigningPayload{
				{
					Address: "addr1",
				},
				{
					Address: "addr2",
				},
				{
					Address: "addr2",
				},
			},
			observed: []string{
				"addr1",
				"addr2",
				"addr2",
			},
			err: true,
		},
		"missing observed signer": {
			intent: []*types.SigningPayload{
				{
					Address: "addr1",
				},
				{
					Address: "addr2",
				},
				{
					Address: "addr2",
				},
			},
			observed: []string{
				"addr1",
			},
			err: true,
		},
		"extra observed signer": {
			intent: []*types.SigningPayload{
				{
					Address: "addr1",
				},
			},
			observed: []string{
				"addr1",
				"addr2",
			},
			err: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := ExpectedSigners(test.intent, test.observed)
			if test.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
