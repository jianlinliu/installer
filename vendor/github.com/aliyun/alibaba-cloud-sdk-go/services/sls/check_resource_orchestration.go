package sls

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

// CheckResourceOrchestration invokes the sls.CheckResourceOrchestration API synchronously
func (client *Client) CheckResourceOrchestration(request *CheckResourceOrchestrationRequest) (response *CheckResourceOrchestrationResponse, err error) {
	response = CreateCheckResourceOrchestrationResponse()
	err = client.DoAction(request, response)
	return
}

// CheckResourceOrchestrationWithChan invokes the sls.CheckResourceOrchestration API asynchronously
func (client *Client) CheckResourceOrchestrationWithChan(request *CheckResourceOrchestrationRequest) (<-chan *CheckResourceOrchestrationResponse, <-chan error) {
	responseChan := make(chan *CheckResourceOrchestrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckResourceOrchestration(request)
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

// CheckResourceOrchestrationWithCallback invokes the sls.CheckResourceOrchestration API asynchronously
func (client *Client) CheckResourceOrchestrationWithCallback(request *CheckResourceOrchestrationRequest, callback func(response *CheckResourceOrchestrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckResourceOrchestrationResponse
		var err error
		defer close(result)
		response, err = client.CheckResourceOrchestration(request)
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

// CheckResourceOrchestrationRequest is the request struct for api CheckResourceOrchestration
type CheckResourceOrchestrationRequest struct {
	*requests.RpcRequest
	Language        string `position:"Query" name:"Language"`
	OperationPolicy string `position:"Query" name:"OperationPolicy"`
	AssetsId        string `position:"Query" name:"AssetsId"`
}

// CheckResourceOrchestrationResponse is the response struct for api CheckResourceOrchestration
type CheckResourceOrchestrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCheckResourceOrchestrationRequest creates a request to invoke CheckResourceOrchestration API
func CreateCheckResourceOrchestrationRequest() (request *CheckResourceOrchestrationRequest) {
	request = &CheckResourceOrchestrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "CheckResourceOrchestration", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckResourceOrchestrationResponse creates a response to parse from CheckResourceOrchestration response
func CreateCheckResourceOrchestrationResponse() (response *CheckResourceOrchestrationResponse) {
	response = &CheckResourceOrchestrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
