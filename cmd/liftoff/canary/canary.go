package canaryCmd

import (
	"fmt"
	"time"

	"github.com/framequery/liftoff/internal/gcloud"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:           "canary",
	Short:         "🐦  Run a canary rollout",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		project := viper.GetString("project")
		service := viper.GetString("service")
		image := viper.GetString("image")
		regions := viper.GetStringSlice("regions")
		percentages := viper.GetIntSlice("percentages")
		intervals := viper.GetIntSlice("intervals")
		envVars := viper.GetStringSlice("env-vars")
		ingress := viper.GetString("ingress")
		allowUnauthenticated := viper.GetBool("allow-unauthenticated")
		// deploy revisions
		for _, r := range regions {
			if err := gcloud.Deploy(service, image, r, project, ingress, envVars, allowUnauthenticated); err != nil {
				return err
			}
		}

		for i, pct := range percentages {
			if err := gcloud.SplitTrafficAcrossRegions(service, regions, pct, project); err != nil {
				return err
			}

			if i < len(intervals) {
				fmt.Printf("⏱️  Waiting %d seconds\n", intervals[i])
				time.Sleep(time.Duration(intervals[i]) * time.Second)
			}
		}
		fmt.Println("🎉  Canary complete at 100%")
		return nil
	},
}
