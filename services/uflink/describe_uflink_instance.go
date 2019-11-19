package uflink

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type DescribeUFlinkInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 集群资源ID
	InstanceId *string `required:"true"`
}

type DescribeUFlinkInstanceResponse struct {
	response.CommonBase
	Data InstanceDetailData
}

func (c *UFlinkClient) NewUFlinkInstanceRequest() *DescribeUFlinkInstanceRequest {
	req := &DescribeUFlinkInstanceRequest{}
	c.Client.SetupRequest(req)
	req.SetRetryable(true)
	return req
}

func (c *UFlinkClient) DescribeUFlinkInstance(req *DescribeUFlinkInstanceRequest) (*DescribeUFlinkInstanceResponse, error) {
	var err error
	var res DescribeUFlinkInstanceResponse

	err = c.Client.InvokeAction("DescribeUFlinkInstance", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
