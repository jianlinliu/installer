package emr

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

// ModifyFlow invokes the emr.ModifyFlow API synchronously
func (client *Client) ModifyFlow(request *ModifyFlowRequest) (response *ModifyFlowResponse, err error) {
	response = CreateModifyFlowResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFlowWithChan invokes the emr.ModifyFlow API asynchronously
func (client *Client) ModifyFlowWithChan(request *ModifyFlowRequest) (<-chan *ModifyFlowResponse, <-chan error) {
	responseChan := make(chan *ModifyFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFlow(request)
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

// ModifyFlowWithCallback invokes the emr.ModifyFlow API asynchronously
func (client *Client) ModifyFlowWithCallback(request *ModifyFlowRequest, callback func(response *ModifyFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFlowResponse
		var err error
		defer close(result)
		response, err = client.ModifyFlow(request)
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

// ModifyFlowRequest is the request struct for api ModifyFlow
type ModifyFlowRequest struct {
	*requests.RpcRequest
	CronExpr                string           `position:"Query" name:"CronExpr"`
	Periodic                requests.Boolean `position:"Query" name:"Periodic"`
	Description             string           `position:"Query" name:"Description"`
	AlertUserGroupBizId     string           `position:"Query" name:"AlertUserGroupBizId"`
	Lifecycle               string           `position:"Query" name:"Lifecycle"`
	HostName                string           `position:"Query" name:"HostName"`
	CreateCluster           requests.Boolean `position:"Query" name:"CreateCluster"`
	EndSchedule             requests.Integer `position:"Query" name:"EndSchedule"`
	Id                      string           `position:"Query" name:"Id"`
	AlertConf               string           `position:"Query" name:"AlertConf"`
	ProjectId               string           `position:"Query" name:"ProjectId"`
	ParentFlowList          string           `position:"Query" name:"ParentFlowList"`
	LogArchiveLocation      string           `position:"Query" name:"LogArchiveLocation"`
	AlertDingDingGroupBizId string           `position:"Query" name:"AlertDingDingGroupBizId"`
	StartSchedule           requests.Integer `position:"Query" name:"StartSchedule"`
	ClusterId               string           `position:"Query" name:"ClusterId"`
	Application             string           `position:"Query" name:"Application"`
	Name                    string           `position:"Query" name:"Name"`
	Namespace               string           `position:"Query" name:"Namespace"`
	Status                  string           `position:"Query" name:"Status"`
	ParentCategory          string           `position:"Query" name:"ParentCategory"`
}

// ModifyFlowResponse is the response struct for api ModifyFlow
type ModifyFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateModifyFlowRequest creates a request to invoke ModifyFlow API
func CreateModifyFlowRequest() (request *ModifyFlowRequest) {
	request = &ModifyFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlow", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyFlowResponse creates a response to parse from ModifyFlow response
func CreateModifyFlowResponse() (response *ModifyFlowResponse) {
	response = &ModifyFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
