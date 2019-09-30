package uflink

import "github.com/ucloud/ucloud-sdk-go/ucloud/request"

type CreateUFlinkApplicationRequest struct {
	request.CommonBase

	InstanceId *string `required:"true"`

	Name *string `required:"true"`

	FlinkVersion *string `required:"true"`

}