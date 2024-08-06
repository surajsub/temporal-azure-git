package models

type VNetApplyOutput struct {
	VNETID   string `json:"vnet_id"`
	VNETGUID string `json:"vnet_guid_id"`
	VNETNAME string `json:"vnet_name"`
}

type VNetCommonOutput struct {
	Value string `json:"value"`
}
