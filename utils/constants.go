package utils

import (
	"fmt"
	"time"
)

const WORKFLOW_TASK_QUEUE = "AZ_STACK_QUEUE"
const BASETFDIRECTORY = "./terraform"

const TEMPORAL_QUEUE_NAME = "provision-task-queue"

// Terraform file locations

const RG_DIR = "/rg"
const RG_ID = "rg_id"
const RG_NAME = "rg_name"

const VNET_DIR = "/vnet"
const VPCTIMEOUT = 1 * time.Minute

const AKS_ENV = "env=%s"

const VNET_ID = "vnet_id"
const VNET_GUID_ID = "vnet_guid_id"
const VNET_NAME = "vnet_name"
const VNET_INFO = "vnet_name=%s"

const subnet_name = "subnet_name"

const SUBNET_DIR = "/subnet"
const SubnetTimeOut = 2 * time.Minute
const SUBNET_1 = "subnet_id_1"
const SUBNET_2 = "subnet_id_2"
const AKS_SUBNET_ID = "aks_subnet_id"
const AKS_APP_SUBNET_ID = "aks_app_subnet_id"
const AKS_SUBNET_NAME = "aks_subnet_name"

const NGINX_DIR = "/nginx"
const DNS_DIR = "/dns"
const DNS_ID = "dns_id"
const PIP_DIR = "/publicip"
const PIP_ID = "public_ip_id"
const PIP_IP = "public_ip"

const NP_DIR = "/nodepool"
const MI_DIR = "/identity"
const MI_ID = "mi_id"
const MI_CLIENT_ID = "mi_client_id"
const MI_PRINCIPAL_ID = "mi_principal_id"
const MI_TENANT_ID = "mi_tenant_id"

const AKS_DIR = "/aks"
const AKS_ID = "aks_id"
const IGW_DIR = "/igw"
const IgwTimeOut = 5 * time.Minute

const NAT_DIR = "/nat"
const RT_DIR = "/route_table"
const SG_DIR = "/sg"
const EC2_DIR = "/ec2"
const EKS_DIR = "/eks"

var TF_INIT = fmt.Sprintf("terraform", "init", "-input=false")

const INIT_ACTIVITY = "Calling the init activity with .."
const APPLY_ACTIVITY = "Calling the apply activity with .."
const OUTPUT_ACTIVITY = "Calling the output activity with .."
const DESTROY_ACTIVITY = "Calling the destroy activity with .."
const GIT_CREATE_TICKET = "Calling the GIT Create Ticket Activity"
const GIT_POLL_TICKET = "Calling the GIT Poll for Status Activity"

const TF_INIT_FAILED = "Failed to execute the terraform init command"
const TF_APPLY_FAILED = "Failed to execute the terraform apply command"
const TF_OUTPUT_FAILED = "Failed to execute the terraform output command"

const RGWorkflow = "Azure Resource Group"
const VNetWorkflow = "Azure Virtual Network"
const NodePoolWorkflow = "Azure Node Pool"
const ManagedIdentityWorkflow = "Azure Managed Identity"
const DeployAppWorkflow = "Azure Application Deployment Workflow"

const SubnetWorkflowAZ = "Azure Virtual Subnet"
const GitWorkflow = "Git Workflow"
const AKSWorkflow = "Azure Kubernetes Service"
const VpcWorkflow = "AWS VPC"
const IGW_WorkflowName = "AWS_Internet_Gateway"
const SubnetWorkflow = "AWS VPC Subnet"
const NatWorkflow = "AWS Nat Service"
const RtWorkflow = "AWS Route Table Service"
const SgWorkflow = "AWS Security Group"
const Ec2Workflow = "AWS EC2 Instance"

// Define the constants for the variables

const VPC_INIT = "Starting the VPC Init Activity:"
const SUBNET_INIT = "Subnet Init Activity:"
const IGW_INIT = "Internet Gateway Init Activity:"
const NAT_INIT = "NAT Init Activity"
const RT_INIT = "Route Table Init Activity"
const SG_INIT = "Security Group Init Activity:"
const EC2_INIT = "EC2 Init Activity:"
const EKS_INIT = "EKS Init Activity"
const NODE_INIT = "EKS Node Init Activity"

const VPC_APPLY = "VPC Apply Activity:"
const SG_APPLY = "Security Group Apply Activity:"
const SUBNET_APPLY = "AWS Subnet Apply Activity"
const NAT_APPLY = "NAT Apply Activity"
const RT_APPLY = "Route Table Apply Activity"
const EC2_APPLY = "EC2 Apply Activity:"
const EKS_APPLY = "EKS Apply Activity:"
const NODE_APPLY = "EKS Node Apply Activity:"

const VPCID = "vpc_id"
const VPCCIDR = "vpc_cidr_block"

const SUBNETID = "subnet_id"
const SUBNETARN = "subnet_arn"

const PRIVATE_SUBNET_ID = "private_subnet_id"
const PUBLIC_SUBNET_ID = "public_subnet_id"

const IGWID = "igw_id"
const IGWARN = "igw_arn"

const SGID = "sg_id"
const SGARN = "sg_arn"

const NATID = "nat_id"
const NATGATEWAYID = "nat_gateway_id"
const NATALLOCATIONID = "nat_allocation_id"

// These need to map to the output from the tf file
const EC2ID = "instance_id"
const EC2PUBLIC = "instance_public_ip"

const EKS_ID = "eks_id"
const EKS_ARN = "eks_arn"
const EKS_ENDPOINT = "eks_endpoint"
