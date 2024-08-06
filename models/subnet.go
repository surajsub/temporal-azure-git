package models

type SubnetCommonOutput struct {
	Value string `json:"value"`
}

type SubnetApplyOutput struct {
	AKSSubnetID    string `json:"aks_subnet_id"`
	AKSAppSubnetID string `json:"aks_app_subnet_id"`
	AKSubnetName   string `json:"aks_subnet_name"`
}
