package cr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetRepoBuildRuleList invokes the cr.GetRepoBuildRuleList API synchronously
func (client *Client) GetRepoBuildRuleList(request *GetRepoBuildRuleListRequest) (response *GetRepoBuildRuleListResponse, err error) {
	response = CreateGetRepoBuildRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoBuildRuleListWithChan invokes the cr.GetRepoBuildRuleList API asynchronously
func (client *Client) GetRepoBuildRuleListWithChan(request *GetRepoBuildRuleListRequest) (<-chan *GetRepoBuildRuleListResponse, <-chan error) {
	responseChan := make(chan *GetRepoBuildRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoBuildRuleList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetRepoBuildRuleListWithCallback invokes the cr.GetRepoBuildRuleList API asynchronously
func (client *Client) GetRepoBuildRuleListWithCallback(request *GetRepoBuildRuleListRequest, callback func(response *GetRepoBuildRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoBuildRuleListResponse
		var err error
		defer close(result)
		response, err = client.GetRepoBuildRuleList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetRepoBuildRuleListRequest is the request struct for api GetRepoBuildRuleList
type GetRepoBuildRuleListRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// GetRepoBuildRuleListResponse is the response struct for api GetRepoBuildRuleList
type GetRepoBuildRuleListResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoBuildRuleListRequest creates a request to invoke GetRepoBuildRuleList API
func CreateGetRepoBuildRuleListRequest() (request *GetRepoBuildRuleListRequest) {
	request = &GetRepoBuildRuleListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildRuleList", "/repos/[RepoNamespace]/[RepoName]/rules", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoBuildRuleListResponse creates a response to parse from GetRepoBuildRuleList response
func CreateGetRepoBuildRuleListResponse() (response *GetRepoBuildRuleListResponse) {
	response = &GetRepoBuildRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
