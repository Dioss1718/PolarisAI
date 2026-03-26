package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
	pluginpkg "github.com/diya-suryawanshi/cloud/backend/plugin"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}

	pluginpkg.GitOps = &gitops.Plugin{}

	if err := orchestrator.Run(); err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}
}
