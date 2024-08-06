package models

type BuildInstanceRequest struct {
	Attributes map[string]string `json:"attributes"`
}

type BuildInstanceResponse struct {
	RequestID string `json:"request_id"`
}

type GitData struct {
	Owner    string
	RepoName string
	GitToken string
}

type AKS struct {
	SubscriptionID        string
	ResourceGroup         string
	Location              string
	Env                   string
	VnetName              string
	AKSClusterName        string
	AKSVersion            string
	AKSVmSize             string
	AKSDnsName            string
	AKSNodePoolName       string
	AKSKubeConfigFileName string
	AKSAppName            string
}

type DNSCommonOutput struct {
	Value string `json:"value"`
}

type DNSApplyOutput struct {
	DNSFQDN string `json:"dns_fqdn"`
	DNSID   string `json:"dns_id"`
	DNSAID  string `json:"dns_a_id"`
}

type PIPCommonOutput struct {
	Value string `json:"value"`
}

type PIPApplyOutput struct {
	PublicIP   string `json:"public_ip"`
	PublicIPID string `json:"public_ip_id"`
}

type NPCommonOutput struct {
	Value string `json:"value"`
}

type NPApplyOutput struct {
	//NP Apply Output Structure.
	InstanceID       string `json:"instance_id"`
	InstancePublicIP string `json:"instance_public_ip"`
}

type AKSCommonOutput struct {
	Value string `json:"value"`
}

type AKSApplyOutput struct {
	//AKS Apply Output Structure.
	ClientCertificate    string `json:"client_certificate"`
	KubeConfig           string `json:"kube_config"`
	KubeHost             string `json:"kube_host"`
	ClusterCaCertificate string `json:"cluster_ca_certificate"`
}

type MICommonOutput struct {
	Value string `json:"value"`
}
type MIApplyOutput struct {
	//MI Apply Output Structure.
	MIID          string `json:"mi_id"`
	MIClientID    string `json:"mi_client_id"`
	MIPrincipalID string `json:"mi_principal_id"`
	MITenantID    string `json:"mi_tenant_id"`
}
