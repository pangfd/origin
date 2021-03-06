package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicasClient is the client for the Replicas methods of the Servicefabric service.
type ReplicasClient struct {
	BaseClient
}

// NewReplicasClient creates an instance of the ReplicasClient client.
func NewReplicasClient(timeout *int32) ReplicasClient {
	return NewReplicasClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewReplicasClientWithBaseURI creates an instance of the ReplicasClient client.
func NewReplicasClientWithBaseURI(baseURI string, timeout *int32) ReplicasClient {
	return ReplicasClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get replicas
// Parameters:
// partitionID - the id of the partition
// replicaID - the id of the replica
func (client ReplicasClient) Get(ctx context.Context, partitionID string, replicaID string) (result Replica, err error) {
	req, err := client.GetPreparer(ctx, partitionID, replicaID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicasClient) GetPreparer(ctx context.Context, partitionID string, replicaID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partitionId": autorest.Encode("path", partitionID),
		"replicaId":   autorest.Encode("path", replicaID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Partitions/{partitionId}/$/GetReplicas/{replicaId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicasClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicasClient) GetResponder(resp *http.Response) (result Replica, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list replicas
// Parameters:
// partitionID - the id of the partition
func (client ReplicasClient) List(ctx context.Context, partitionID string) (result ReplicaList, err error) {
	req, err := client.ListPreparer(ctx, partitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ReplicasClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicasClient) ListPreparer(ctx context.Context, partitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partitionId": autorest.Encode("path", partitionID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Partitions/{partitionId}/$/GetReplicas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicasClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicasClient) ListResponder(resp *http.Response) (result ReplicaList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
