package uflink

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type GetUFlinkSubmittedJobLogRequest struct {
	request.CommonBase

	InstanceId *string `required:"true"`

	Id *int `required:"true"`
}

type GetUFlinkSubmittedJobLogResponse struct {
	response.CommonBase

	Data string
}

func (c *UFlinkClient) NewGetUFlinkSubmittedJobLogRequest() *GetUFlinkSubmittedJobLogRequest {
	req := &GetUFlinkSubmittedJobLogRequest{}
	c.Client.SetupRequest(req)
	req.SetRetryable(true)
	return req
}