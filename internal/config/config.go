package config

import (
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
	return nil
}

// BindFlags attaches CLI flags to viper keys and writes config on exit
func BindFlags(cmd *cobra.Command) {
	f := cmd.PersistentFlags()
	f.StringP("project", "p", "", "GCP project ID")
	f.StringP("service", "s", "", "Cloud Run service name")
	f.StringP("image", "i", "", "Container image URL for canary")
	f.StringSlice("regions", nil, "GCP regions (comma-separated)")
	f.IntSlice("percentages", nil, "Traffic percentages e.g. 10,50,100")
	f.IntSlice("intervals", nil, "Intervals (s) between steps")
	f.StringSlice("env-vars", nil, "Environment variables (KEY=VALUE) to set on each revision")

	viper.BindPFlag("env-vars", f.Lookup("env-vars"))
	viper.BindPFlag("project", f.Lookup("project"))
	viper.BindPFlag("service", f.Lookup("service"))
	viper.BindPFlag("image", f.Lookup("image"))
	viper.BindPFlag("regions", f.Lookup("regions"))
	viper.BindPFlag("percentages", f.Lookup("percentages"))
	viper.BindPFlag("intervals", f.Lookup("intervals"))

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// persist config
		tmp := os.TempDir()
		return viper.WriteConfigAs(filepath.Join(tmp, "liftoff_config.json"))
	}
}
