package storage

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// FileSharesClient is the the Azure Storage Management API.
type FileSharesClient struct {
	BaseClient
}

// NewFileSharesClient creates an instance of the FileSharesClient client.
func NewFileSharesClient(subscriptionID string) FileSharesClient {
	return NewFileSharesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFileSharesClientWithBaseURI creates an instance of the FileSharesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFileSharesClientWithBaseURI(baseURI string, subscriptionID string) FileSharesClient {
	return FileSharesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new share under the specified account as described by request body. The share resource includes
// metadata and properties for that share. It does not include a list of the files contained by the share.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// shareName - the name of the file share within the specified storage account. File share names must be
// between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
// character must be immediately preceded and followed by a letter or number.
// fileShare - properties of the file share to create.
func (client FileSharesClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare) (result FileShare, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: shareName,
			Constraints: []validation.Constraint{{Target: "shareName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "shareName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: fileShare,
			Constraints: []validation.Constraint{{Target: "fileShare.FileShareProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "fileShare.FileShareProperties.ShareQuota", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "fileShare.FileShareProperties.ShareQuota", Name: validation.InclusiveMaximum, Rule: int64(102400), Chain: nil},
						{Target: "fileShare.FileShareProperties.ShareQuota", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
					}},
				}}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareName, fileShare)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client FileSharesClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", pathParameters),
		autorest.WithJSON(fileShare),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client FileSharesClient) CreateResponder(resp *http.Response) (result FileShare, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes specified share under its account.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// shareName - the name of the file share within the specified storage account. File share names must be
// between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
// character must be immediately preceded and followed by a letter or number.
func (client FileSharesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: shareName,
			Constraints: []validation.Constraint{{Target: "shareName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "shareName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FileSharesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FileSharesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets properties of a specified share.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// shareName - the name of the file share within the specified storage account. File share names must be
// between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
// character must be immediately preceded and followed by a letter or number.
// expand - optional, used to expand the properties within share's properties.
func (client FileSharesClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, expand GetShareExpand) (result FileShare, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: shareName,
			Constraints: []validation.Constraint{{Target: "shareName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "shareName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FileSharesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, expand GetShareExpand) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FileSharesClient) GetResponder(resp *http.Response) (result FileShare, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all shares.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// maxpagesize - optional. Specified maximum number of shares that can be included in the list.
// filter - optional. When specified, only share names starting with the filter will be listed.
// expand - optional, used to expand the properties within share's properties.
func (client FileSharesClient) List(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, expand ListSharesExpand) (result FileShareItemsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.List")
		defer func() {
			sc := -1
			if result.fsi.Response.Response != nil {
				sc = result.fsi.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, maxpagesize, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fsi.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "List", resp, "Failure sending request")
		return
	}

	result.fsi, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "List", resp, "Failure responding to request")
	}
	if result.fsi.hasNextLink() && result.fsi.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client FileSharesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, expand ListSharesExpand) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(maxpagesize) > 0 {
		queryParameters["$maxpagesize"] = autorest.Encode("query", maxpagesize)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FileSharesClient) ListResponder(resp *http.Response) (result FileShareItems, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client FileSharesClient) listNextResults(ctx context.Context, lastResults FileShareItems) (result FileShareItems, err error) {
	req, err := lastResults.fileShareItemsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storage.FileSharesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storage.FileSharesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FileSharesClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, maxpagesize string, filter string, expand ListSharesExpand) (result FileShareItemsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, maxpagesize, filter, expand)
	return
}

// Restore restore a file share within a valid retention days if share soft delete is enabled
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// shareName - the name of the file share within the specified storage account. File share names must be
// between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
// character must be immediately preceded and followed by a letter or number.
func (client FileSharesClient) Restore(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare DeletedShare) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.Restore")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: shareName,
			Constraints: []validation.Constraint{{Target: "shareName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "shareName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: deletedShare,
			Constraints: []validation.Constraint{{Target: "deletedShare.DeletedShareName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "deletedShare.DeletedShareVersion", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "Restore", err.Error())
	}

	req, err := client.RestorePreparer(ctx, resourceGroupName, accountName, shareName, deletedShare)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Restore", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestoreSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Restore", resp, "Failure sending request")
		return
	}

	result, err = client.RestoreResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Restore", resp, "Failure responding to request")
	}

	return
}

// RestorePreparer prepares the Restore request.
func (client FileSharesClient) RestorePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, deletedShare DeletedShare) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}/restore", pathParameters),
		autorest.WithJSON(deletedShare),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestoreSender sends the Restore request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) RestoreSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RestoreResponder handles the response to the Restore request. The method always
// closes the http.Response Body.
func (client FileSharesClient) RestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates share properties as specified in request body. Properties not mentioned in the request will not be
// changed. Update fails if the specified share does not already exist.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// shareName - the name of the file share within the specified storage account. File share names must be
// between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-)
// character must be immediately preceded and followed by a letter or number.
// fileShare - properties to update for the file share.
func (client FileSharesClient) Update(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare) (result FileShare, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileSharesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: shareName,
			Constraints: []validation.Constraint{{Target: "shareName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "shareName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.FileSharesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, shareName, fileShare)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.FileSharesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FileSharesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, fileShare FileShare) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", pathParameters),
		autorest.WithJSON(fileShare),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FileSharesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FileSharesClient) UpdateResponder(resp *http.Response) (result FileShare, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
