package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	cli "github.com/argoproj/argocd-cli/cmd/argocd"
	util "github.com/argoproj/argocd-cli/cmd/argocdutil"
)

const (
	binaryNameEnv = "ARGOCD_BINARY_NAME"
)

func main() {
	var command *cobra.Command

	binaryName := filepath.Base(os.Args[0])
	if val := os.Getenv(binaryNameEnv); val != "" {
		binaryName = val
	}
	switch binaryName {
	case "argocd", "argocd-linux-amd64", "argocd-darwin-amd64", "argocd-windows-amd64.exe":
		command = cli.NewCommand()
	case "argocd-util", "argocd-util-linux-amd64", "argocd-util-darwin-amd64", "argocd-util-windows-amd64.exe":
		command = util.NewCommand()

	}

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
