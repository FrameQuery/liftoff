package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds CLI defaults
var DefaultConfig = struct {
	Project     string
	Service     string
	Image       string
	Regions     []string
	Percentages []int
	Intervals   []int
}{}

// Load reads config from temp JSON or env
func Load() error {
	viper.SetConfigType("json")
	tmp := os.TempDir()
	viper.SetConfigFile(filepath.Join(tmp, "liftoff_config.json"))
	if err := viper.ReadInConfig(); err == nil {
		return nil
	}
	// no config, use defaults
	viper.SetDefault("regions", []string{"europe-west2", "europe-west4"})
	viper.SetDefault("percentages", []int{10, 50, 100})
	viper.SetDefault("intervals", []int{300, 300})
	viper.SetDefault("ingress", "internal-only")
	return nil
}

// BindFlags attaches CLI flags to viper keys and writes config on exit
func BindFlags(cmd *cobra.Command) {
	f := cmd.PersistentFlags()
	f.StringP("project", "p", "", "GCP project ID")
	f.StringP("service", "s", "", "Cloud Run service name")
	f.StringP("image", "i", "", "Container image URL for canary")
	f.StringP("ingress", "in", "", "Ingress settings (all, internal-only, internal-and-cloud-run)")
	f.StringSlice("regions", nil, "GCP regions (comma-separated)")
	f.IntSlice("percentages", nil, "Traffic percentages e.g. 10,50,100")
	f.IntSlice("intervals", nil, "Intervals (s) between steps")
	f.StringSlice("env-vars", nil, "Environment variables (KEY=VALUE) to set on each revision")

	if err := viper.BindPFlag("env-vars", f.Lookup("env-vars")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("project", f.Lookup("project")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("service", f.Lookup("service")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("image", f.Lookup("image")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("regions", f.Lookup("regions")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("percentages", f.Lookup("percentages")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("percentages", f.Lookup("percentages")); err != nil {

	}
	if err := viper.BindPFlag("intervals", f.Lookup("intervals")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}
	if err := viper.BindPFlag("ingress", f.Lookup("ingress")); err != nil {
		fmt.Printf("⚠️  ERROR: %v\n", err)
	}

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// persist config
		tmp := os.TempDir()
		return viper.WriteConfigAs(filepath.Join(tmp, "liftoff_config.json"))
	}
}
