package cr_ee

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

// CreateRepoSyncTaskByRule invokes the cr.CreateRepoSyncTaskByRule API synchronously
// api document: https://help.aliyun.com/api/cr/createreposynctaskbyrule.html
func (client *Client) CreateRepoSyncTaskByRule(request *CreateRepoSyncTaskByRuleRequest) (response *CreateRepoSyncTaskByRuleResponse, err error) {
	response = CreateCreateRepoSyncTaskByRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepoSyncTaskByRuleWithChan invokes the cr.CreateRepoSyncTaskByRule API asynchronously
// api document: https://help.aliyun.com/api/cr/createreposynctaskbyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoSyncTaskByRuleWithChan(request *CreateRepoSyncTaskByRuleRequest) (<-chan *CreateRepoSyncTaskByRuleResponse, <-chan error) {
	responseChan := make(chan *CreateRepoSyncTaskByRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepoSyncTaskByRule(request)
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

// CreateRepoSyncTaskByRuleWithCallback invokes the cr.CreateRepoSyncTaskByRule API asynchronously
// api document: https://help.aliyun.com/api/cr/createreposynctaskbyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoSyncTaskByRuleWithCallback(request *CreateRepoSyncTaskByRuleRequest, callback func(response *CreateRepoSyncTaskByRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepoSyncTaskByRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateRepoSyncTaskByRule(request)
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

// CreateRepoSyncTaskByRuleRequest is the request struct for api CreateRepoSyncTaskByRule
type CreateRepoSyncTaskByRuleRequest struct {
	*requests.RpcRequest
	RepoId     string `position:"Query" name:"RepoId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Tag        string `position:"Query" name:"Tag"`
	SyncRuleId string `position:"Query" name:"SyncRuleId"`
}

// CreateRepoSyncTaskByRuleResponse is the response struct for api CreateRepoSyncTaskByRule
type CreateRepoSyncTaskByRuleResponse struct {
	*responses.BaseResponse
	CreateRepoSyncTaskByRuleIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                              string `json:"Code" xml:"Code"`
	RequestId                         string `json:"RequestId" xml:"RequestId"`
	SyncTaskId                        string `json:"SyncTaskId" xml:"SyncTaskId"`
}

// CreateCreateRepoSyncTaskByRuleRequest creates a request to invoke CreateRepoSyncTaskByRule API
func CreateCreateRepoSyncTaskByRuleRequest() (request *CreateRepoSyncTaskByRuleRequest) {
	request = &CreateRepoSyncTaskByRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "CreateRepoSyncTaskByRule", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRepoSyncTaskByRuleResponse creates a response to parse from CreateRepoSyncTaskByRule response
func CreateCreateRepoSyncTaskByRuleResponse() (response *CreateRepoSyncTaskByRuleResponse) {
	response = &CreateRepoSyncTaskByRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
