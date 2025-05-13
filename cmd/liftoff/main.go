package main

import (
	"fmt"
	"os"

	canaryCmd "github.com/framequery/liftoff/cmd/liftoff/canary"
	configcmd "github.com/framequery/liftoff/cmd/liftoff/config"
	versionCmd "github.com/framequery/liftoff/cmd/liftoff/version"
	"github.com/framequery/liftoff/internal/auth"
	"github.com/framequery/liftoff/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "liftoff",
	Short: "🚀 Liftoff: Multi-region Cloud Run canary deploy",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// load config
		if err := config.Load(); err != nil {
			return err
		}
		// ensure auth
		return auth.Ensure()
	},
}

func init() {
	// bind flags to viper
	config.BindFlags(rootCmd)

	rootCmd.AddCommand(canaryCmd.Cmd)
	rootCmd.AddCommand(configcmd.Cmd)
	rootCmd.AddCommand(versionCmd.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
		os.Exit(1)
	}
}
