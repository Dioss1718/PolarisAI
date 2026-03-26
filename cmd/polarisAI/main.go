package main

import (
	"log"

	"github.com/joho/godotenv"

	pluginpkg "github.com/diya-suryawanshi/cloud/backend/plugin"
	"github.com/diya-suryawanshi/cloud/backend/server"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}

	pluginpkg.GitOps = &gitops.Plugin{}

	log.Println("Starting Polaris Autonomous Cloud Governance API on :8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
