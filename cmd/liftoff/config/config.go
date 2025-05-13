package configcmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "⚙️ Manage liftoff defaults",
}

func init() {
	Cmd.AddCommand(setCmd)
	Cmd.AddCommand(viewCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [service]",
	Short: "💾  Set default config for a service",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		service := args[0]
		// read existing
		tmp := os.TempDir()
		file := filepath.Join(tmp, "liftoff_config.json")
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			// start fresh
			viper.Set("services", map[string]interface{}{})
		}
		// capture flags
		cfg := make(map[string]interface{})
		for _, key := range []string{"project", "image", "regions", "percentages", "intervals"} {
			if viper.IsSet(key) {
				cfg[key] = viper.Get(key)
			}
		}
		services := viper.GetStringMap("services")
		services[service] = cfg
		viper.Set("services", services)
		if err := viper.WriteConfigAs(file); err != nil {
			return err
		}
		fmt.Printf("✅  Defaults saved for service '%s'\n", service)
		return nil
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "🔍  View all service configs",
	RunE: func(cmd *cobra.Command, args []string) error {
		tmp := os.TempDir()
		file := filepath.Join(tmp, "liftoff_config.json")
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("No config found.")
			return nil
		}
		fmt.Println(viper.Get("services"))
		return nil
	},
}
