# temporal-azure-git
How to deploy AKS + App in Azure using temporal and Terraform. 
 Temporal AKS Deployment with Terraform

This repository provides a guide to deploying an Azure Kubernetes Service (AKS) cluster and a sample pet store application using Terraform and Temporal workflows.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Terraform](https://www.terraform.io/downloads.html)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)
- [Temporal CLI](https://docs.temporal.io/docs/getting-started)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)

## Setup

### 1. Configure Azure CLI

Ensure you are logged in to your Azure account:
az login
az account show - if you have multiple accounts configured then you must pick the right subscriptionid to be passed in the .env file. 

### 2. Make sure you have kubectl installed on your machine
kubectl version should return a value - Mine returns v1.28.2

### 3. Make sure you have terraform installed on your machine
