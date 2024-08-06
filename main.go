package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/surajsub/temporal-azure-git/models"
	"github.com/surajsub/temporal-azure-git/worker"
	"github.com/surajsub/temporal-azure-git/workflows"
	"log"
	"os"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}

	azureSubscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if azureSubscriptionID == "" {
		log.Fatal("AZURE_SUBSCRIPTION_ID environment variable not set")
	}

	rgname := os.Getenv("RGNAME")
	if rgname == "" {
		log.Fatal("RGNAME environment variable not set")
	}

	switch os.Args[1] {
	case "worker":
		worker.RunWorker()
	case "starter":

		starterCmd := flag.NewFlagSet("starter", flag.ExitOnError)
		vpcCdirBlock := starterCmd.String("vpcCdirBlock", "", "CIDR Block for the VPC")

		err := starterCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("failed to parse 'starter' flags")
		}

		// Ensure the required flags are provided
		if *vpcCdirBlock == "" {
			log.Fatal("VPC Cdir block is not provided")
		}

		// this data should come from the command line as input but for now.. just setting it like that
		workflows.StartWorkflow(*vpcCdirBlock, models.AKS{
			ResourceGroup:   rgname,
			Location:        os.Getenv("LOCATION"),
			Env:             os.Getenv("ENV"),
			VnetName:        os.Getenv("VNETNAME"),
			AKSClusterName:  os.Getenv("AKSClusterName"),
			AKSVersion:      os.Getenv("AKSVersion"),
			AKSVmSize:       os.Getenv("AKSVmSize"),
			AKSDnsName:      os.Getenv("AKSDnsName"),
			AKSNodePoolName: os.Getenv("AKSNodePoolName"),
			SubscriptionID:  azureSubscriptionID,
		}, models.GitData{
			Owner:    "REPO-OWNER", // Adjust these variables to your specific values
			RepoName: "REPO-NAME",
			GitToken: "github_pat_YOUR_GITHUBTOKEN",
		})

		// workflows.StartWorkflow()

	default:
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}
}
