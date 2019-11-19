package uflink

type InstanceDetailData struct {
	Cluster       InstanceData
	ProxyUrl      string
	AdminUserName string
}

type InstanceData struct {
	InstanceName string
	State        string
	InstanceId   string
}
