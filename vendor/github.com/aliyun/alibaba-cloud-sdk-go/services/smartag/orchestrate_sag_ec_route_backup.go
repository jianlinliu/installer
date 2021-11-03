package smartag

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

// OrchestrateSagECRouteBackup invokes the smartag.OrchestrateSagECRouteBackup API synchronously
func (client *Client) OrchestrateSagECRouteBackup(request *OrchestrateSagECRouteBackupRequest) (response *OrchestrateSagECRouteBackupResponse, err error) {
	response = CreateOrchestrateSagECRouteBackupResponse()
	err = client.DoAction(request, response)
	return
}

// OrchestrateSagECRouteBackupWithChan invokes the smartag.OrchestrateSagECRouteBackup API asynchronously
func (client *Client) OrchestrateSagECRouteBackupWithChan(request *OrchestrateSagECRouteBackupRequest) (<-chan *OrchestrateSagECRouteBackupResponse, <-chan error) {
	responseChan := make(chan *OrchestrateSagECRouteBackupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OrchestrateSagECRouteBackup(request)
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

// OrchestrateSagECRouteBackupWithCallback invokes the smartag.OrchestrateSagECRouteBackup API asynchronously
func (client *Client) OrchestrateSagECRouteBackupWithCallback(request *OrchestrateSagECRouteBackupRequest, callback func(response *OrchestrateSagECRouteBackupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OrchestrateSagECRouteBackupResponse
		var err error
		defer close(result)
		response, err = client.OrchestrateSagECRouteBackup(request)
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

// OrchestrateSagECRouteBackupRequest is the request struct for api OrchestrateSagECRouteBackup
type OrchestrateSagECRouteBackupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
}

// OrchestrateSagECRouteBackupResponse is the response struct for api OrchestrateSagECRouteBackup
type OrchestrateSagECRouteBackupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOrchestrateSagECRouteBackupRequest creates a request to invoke OrchestrateSagECRouteBackup API
func CreateOrchestrateSagECRouteBackupRequest() (request *OrchestrateSagECRouteBackupRequest) {
	request = &OrchestrateSagECRouteBackupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "OrchestrateSagECRouteBackup", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOrchestrateSagECRouteBackupResponse creates a response to parse from OrchestrateSagECRouteBackup response
func CreateOrchestrateSagECRouteBackupResponse() (response *OrchestrateSagECRouteBackupResponse) {
	response = &OrchestrateSagECRouteBackupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
